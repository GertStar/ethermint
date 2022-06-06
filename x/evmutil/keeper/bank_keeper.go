package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/Gert-Chain/core/x/evmutil/types"
)

const (
	// EvmDenom is the gas denom used by the evm
	EvmDenom = "agert"

	// CosmosDenom is the gas denom used by the gert app
	CosmosDenom = "ugert"
)

// ConversionMultiplier is the conversion multiplier between agert and ugert
var ConversionMultiplier = sdk.NewInt(1_000_000_000_000)

var _ evmtypes.BankKeeper = EvmBankKeeper{}

// EvmBankKeeper is a BankKeeper wrapper for the x/evm module to allow the use
// of the 18 decimal agert coin on the evm.
// x/evm consumes gas and send coins by minting and burning agert coins in its module
// account and then sending the funds to the target account.
// This keeper uses both the ugert coin and a separate agert balance to manage the
// extra percision needed by the evm.
type EvmBankKeeper struct {
	agertKeeper Keeper
	bk          types.BankKeeper
	ak          types.AccountKeeper
}

func NewEvmBankKeeper(agertKeeper Keeper, bk types.BankKeeper, ak types.AccountKeeper) EvmBankKeeper {
	return EvmBankKeeper{
		agertKeeper: agertKeeper,
		bk:          bk,
		ak:          ak,
	}
}

// GetBalance returns the total **spendable** balance of agert for a given account by address.
func (k EvmBankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	if denom != EvmDenom {
		panic(fmt.Errorf("only evm denom %s is supported by EvmBankKeeper", EvmDenom))
	}

	spendableCoins := k.bk.SpendableCoins(ctx, addr)
	ugert := spendableCoins.AmountOf(CosmosDenom)
	agert := k.agertKeeper.GetBalance(ctx, addr)
	total := ugert.Mul(ConversionMultiplier).Add(agert)
	return sdk.NewCoin(EvmDenom, total)
}

// SendCoinsFromModuleToAccount transfers agert coins from a ModuleAccount to an AccAddress.
// It will panic if the module account does not exist. An error is returned if the recipient
// address is black-listed or if sending the tokens fails.
func (k EvmBankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	ugert, agert, err := SplitAgertCoins(amt)
	if err != nil {
		return err
	}

	if ugert.Amount.IsPositive() {
		if err := k.bk.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, sdk.NewCoins(ugert)); err != nil {
			return err
		}
	}

	senderAddr := k.GetModuleAddress(senderModule)
	if err := k.ConvertOneUgertToAgertIfNeeded(ctx, senderAddr, agert); err != nil {
		return err
	}

	if err := k.agertKeeper.SendBalance(ctx, senderAddr, recipientAddr, agert); err != nil {
		return err
	}

	return k.ConvertAgertToUgert(ctx, recipientAddr)
}

// SendCoinsFromAccountToModule transfers agert coins from an AccAddress to a ModuleAccount.
// It will panic if the module account does not exist.
func (k EvmBankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	ugert, agertNeeded, err := SplitAgertCoins(amt)
	if err != nil {
		return err
	}

	if ugert.IsPositive() {
		if err := k.bk.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, sdk.NewCoins(ugert)); err != nil {
			return err
		}
	}

	if err := k.ConvertOneUgertToAgertIfNeeded(ctx, senderAddr, agertNeeded); err != nil {
		return err
	}

	recipientAddr := k.GetModuleAddress(recipientModule)
	if err := k.agertKeeper.SendBalance(ctx, senderAddr, recipientAddr, agertNeeded); err != nil {
		return err
	}

	return k.ConvertAgertToUgert(ctx, recipientAddr)
}

// MintCoins mints agert coins by minting the equivalent ugert coins and any remaining agert coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	ugert, agert, err := SplitAgertCoins(amt)
	if err != nil {
		return err
	}

	if ugert.IsPositive() {
		if err := k.bk.MintCoins(ctx, moduleName, sdk.NewCoins(ugert)); err != nil {
			return err
		}
	}

	recipientAddr := k.GetModuleAddress(moduleName)
	if err := k.agertKeeper.AddBalance(ctx, recipientAddr, agert); err != nil {
		return err
	}

	return k.ConvertAgertToUgert(ctx, recipientAddr)
}

// BurnCoins burns agert coins by burning the equivalent ugert coins and any remaining agert coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	ugert, agert, err := SplitAgertCoins(amt)
	if err != nil {
		return err
	}

	if ugert.IsPositive() {
		if err := k.bk.BurnCoins(ctx, moduleName, sdk.NewCoins(ugert)); err != nil {
			return err
		}
	}

	moduleAddr := k.GetModuleAddress(moduleName)
	if err := k.ConvertOneUgertToAgertIfNeeded(ctx, moduleAddr, agert); err != nil {
		return err
	}

	return k.agertKeeper.RemoveBalance(ctx, moduleAddr, agert)
}

// ConvertOneUgertToAgertIfNeeded converts 1 ugert to agert for an address if
// its agert balance is smaller than the agertNeeded amount.
func (k EvmBankKeeper) ConvertOneUgertToAgertIfNeeded(ctx sdk.Context, addr sdk.AccAddress, agertNeeded sdk.Int) error {
	agertBal := k.agertKeeper.GetBalance(ctx, addr)
	if agertBal.GTE(agertNeeded) {
		return nil
	}

	ugertToStore := sdk.NewCoins(sdk.NewCoin(CosmosDenom, sdk.OneInt()))
	if err := k.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, ugertToStore); err != nil {
		return err
	}

	// add 1ugert equivalent of agert to addr
	agertToReceive := ConversionMultiplier
	if err := k.agertKeeper.AddBalance(ctx, addr, agertToReceive); err != nil {
		return err
	}

	return nil
}

// ConvertAgertToUgert converts all available agert to ugert for a given AccAddress.
func (k EvmBankKeeper) ConvertAgertToUgert(ctx sdk.Context, addr sdk.AccAddress) error {
	totalAgert := k.agertKeeper.GetBalance(ctx, addr)
	ugert, _, err := SplitAgertCoins(sdk.NewCoins(sdk.NewCoin(EvmDenom, totalAgert)))
	if err != nil {
		return err
	}

	// do nothing if account does not have enough agert for a single ugert
	ugertToReceive := ugert.Amount
	if !ugertToReceive.IsPositive() {
		return nil
	}

	// remove agert used for converting to ugert
	agertToBurn := ugertToReceive.Mul(ConversionMultiplier)
	finalBal := totalAgert.Sub(agertToBurn)
	if err := k.agertKeeper.SetBalance(ctx, addr, finalBal); err != nil {
		return err
	}

	fromAddr := k.GetModuleAddress(types.ModuleName)
	if err := k.bk.SendCoins(ctx, fromAddr, addr, sdk.NewCoins(ugert)); err != nil {
		return err
	}

	return nil
}

func (k EvmBankKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	addr := k.ak.GetModuleAddress(moduleName)
	if addr == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}
	return addr
}

// SplitAgertCoins splits agert coins to the equivalent ugert coins and any remaining agert balance.
// An error will be returned if the coins are not valid or if the coins are not the agert denom.
func SplitAgertCoins(coins sdk.Coins) (sdk.Coin, sdk.Int, error) {
	agert := sdk.ZeroInt()
	ugert := sdk.NewCoin(CosmosDenom, sdk.ZeroInt())

	if len(coins) == 0 {
		return ugert, agert, nil
	}

	if err := ValidateEvmCoins(coins); err != nil {
		return ugert, agert, err
	}

	// note: we should always have len(coins) == 1 here since coins cannot have dup denoms after we validate.
	coin := coins[0]
	remainingBalance := coin.Amount.Mod(ConversionMultiplier)
	if remainingBalance.IsPositive() {
		agert = remainingBalance
	}
	ugertAmount := coin.Amount.Quo(ConversionMultiplier)
	if ugertAmount.IsPositive() {
		ugert = sdk.NewCoin(CosmosDenom, ugertAmount)
	}

	return ugert, agert, nil
}

// ValidateEvmCoins validates the coins from evm is valid and is the EvmDenom (agert).
func ValidateEvmCoins(coins sdk.Coins) error {
	if len(coins) == 0 {
		return nil
	}

	// validate that coins are non-negative, sorted, and no dup denoms
	if err := coins.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, coins.String())
	}

	// validate that coin denom is agert
	if len(coins) != 1 || coins[0].Denom != EvmDenom {
		errMsg := fmt.Sprintf("invalid evm coin denom, only %s is supported", EvmDenom)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errMsg)
	}

	return nil
}

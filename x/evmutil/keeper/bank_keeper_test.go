package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmtime "github.com/tendermint/tendermint/types/time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/Gert-Chain/core/x/evmutil/keeper"
	"github.com/Gert-Chain/core/x/evmutil/testutil"
	"github.com/Gert-Chain/core/x/evmutil/types"
)

type evmBankKeeperTestSuite struct {
	testutil.Suite
}

func (suite *evmBankKeeperTestSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_ReturnsSpendable() {
	startingCoins := sdk.NewCoins(sdk.NewInt64Coin("ugert", 10))
	startingAgert := sdk.NewInt(100)

	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	bacc := authtypes.NewBaseAccountWithAddress(suite.Addrs[0])
	vacc := vesting.NewContinuousVestingAccount(bacc, startingCoins, now.Unix(), endTime.Unix())
	suite.AccountKeeper.SetAccount(suite.Ctx, vacc)

	err := suite.App.FundAccount(suite.Ctx, suite.Addrs[0], startingCoins)
	suite.Require().NoError(err)
	err = suite.Keeper.SetBalance(suite.Ctx, suite.Addrs[0], startingAgert)
	suite.Require().NoError(err)

	coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "agert")
	suite.Require().Equal(startingAgert, coin.Amount)

	ctx := suite.Ctx.WithBlockTime(now.Add(12 * time.Hour))
	coin = suite.EvmBankKeeper.GetBalance(ctx, suite.Addrs[0], "agert")
	suite.Require().Equal(sdk.NewIntFromUint64(5_000_000_000_100), coin.Amount)
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_NotEvmDenom() {
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ugert")
	})
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "busd")
	})
}

func (suite *evmBankKeeperTestSuite) TestGetBalance() {
	tests := []struct {
		name           string
		startingAmount sdk.Coins
		expAmount      sdk.Int
	}{
		{
			"ugert with agert",
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 100),
				sdk.NewInt64Coin("ugert", 10),
			),
			sdk.NewInt(10_000_000_000_100),
		},
		{
			"just agert",
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 100),
				sdk.NewInt64Coin("busd", 100),
			),
			sdk.NewInt(100),
		},
		{
			"just ugert",
			sdk.NewCoins(
				sdk.NewInt64Coin("ugert", 10),
				sdk.NewInt64Coin("busd", 100),
			),
			sdk.NewInt(10_000_000_000_000),
		},
		{
			"no ugert or agert",
			sdk.NewCoins(),
			sdk.ZeroInt(),
		},
		{
			"with avaka that is more than 1 ugert",
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 20_000_000_000_220),
				sdk.NewInt64Coin("ugert", 11),
			),
			sdk.NewInt(31_000_000_000_220),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithgert(suite.Addrs[0], tt.startingAmount)
			coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "agert")
			suite.Require().Equal(tt.expAmount, coin.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromModuleToAccount() {
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin("agert", 200),
		sdk.NewInt64Coin("ugert", 100),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		startingAccBal sdk.Coins
		expAccBal      sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_000_000_000_010)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 10),
				sdk.NewInt64Coin("ugert", 12),
			),
			false,
		},
		{
			"send less than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 122)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 122),
				sdk.NewInt64Coin("ugert", 0),
			),
			false,
		},
		{
			"send an exact amount of ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 98_000_000_000_000)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 0o0),
				sdk.NewInt64Coin("ugert", 98),
			),
			false,
		},
		{
			"send no agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 0)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 0),
				sdk.NewInt64Coin("ugert", 0),
			),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total agert to cover",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough ugert to cover",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts receiver's agert to ugert if there's enough agert after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 99_000_000_000_200)),
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 999_999_999_900),
				sdk.NewInt64Coin("ugert", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 100),
				sdk.NewInt64Coin("ugert", 101),
			),
			false,
		},
		{
			"converts all of receiver's agert to ugert even if somehow receiver has more than 1ugert of agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_000_000_000_100)),
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 5_999_999_999_990),
				sdk.NewInt64Coin("ugert", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 90),
				sdk.NewInt64Coin("ugert", 19),
			),
			false,
		},
		{
			"swap 1 ugert for agert if module account doesn't have enough agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 99_000_000_001_000)),
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 200),
				sdk.NewInt64Coin("ugert", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("agert", 1200),
				sdk.NewInt64Coin("ugert", 100),
			),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithgert(suite.Addrs[0], tt.startingAccBal)
			suite.FundModuleAccountWithgert(evmtypes.ModuleName, startingModuleCoins)

			// fund our module with some ugert to account for converting extra agert back to ugert
			suite.FundModuleAccountWithgert(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("ugert", 10)))

			err := suite.EvmBankKeeper.SendCoinsFromModuleToAccount(suite.Ctx, evmtypes.ModuleName, suite.Addrs[0], tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check ugert
			ugertSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ugert")
			suite.Require().Equal(tt.expAccBal.AmountOf("ugert").Int64(), ugertSender.Amount.Int64())

			// check agert
			actualAgert := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expAccBal.AmountOf("agert").Int64(), actualAgert.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromAccountToModule() {
	startingAccCoins := sdk.NewCoins(
		sdk.NewInt64Coin("agert", 200),
		sdk.NewInt64Coin("ugert", 100),
	)
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin("agert", 100_000_000_000),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		expSenderCoins sdk.Coins
		expModuleCoins sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_000_000_000_010)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 190), sdk.NewInt64Coin("ugert", 88)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_010), sdk.NewInt64Coin("ugert", 12)),
			false,
		},
		{
			"send less than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 122)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 78), sdk.NewInt64Coin("ugert", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_122), sdk.NewInt64Coin("ugert", 0)),
			false,
		},
		{
			"send an exact amount of ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 98_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 200), sdk.NewInt64Coin("ugert", 2)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_000), sdk.NewInt64Coin("ugert", 98)),
			false,
		},
		{
			"send no agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 0)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 200), sdk.NewInt64Coin("ugert", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_000), sdk.NewInt64Coin("ugert", 0)),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("agert", 12_000_000_000_000),
				sdk.NewInt64Coin("agert", 2_000_000_000_000),
			},
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total agert to cover",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough ugert to cover",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts 1 ugert to agert if not enough agert to cover",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 99_001_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 999_000_000_200), sdk.NewInt64Coin("ugert", 0)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 101_000_000_000), sdk.NewInt64Coin("ugert", 99)),
			false,
		},
		{
			"converts receiver's agert to ugert if there's enough agert after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 5_900_000_000_200)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_000_000_000), sdk.NewInt64Coin("ugert", 94)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 200), sdk.NewInt64Coin("ugert", 6)),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundAccountWithgert(suite.Addrs[0], startingAccCoins)
			suite.FundModuleAccountWithgert(evmtypes.ModuleName, startingModuleCoins)

			err := suite.EvmBankKeeper.SendCoinsFromAccountToModule(suite.Ctx, suite.Addrs[0], evmtypes.ModuleName, tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check sender balance
			ugertSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ugert")
			suite.Require().Equal(tt.expSenderCoins.AmountOf("ugert").Int64(), ugertSender.Amount.Int64())
			actualAgert := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expSenderCoins.AmountOf("agert").Int64(), actualAgert.Int64())

			// check module balance
			moduleAddr := suite.AccountKeeper.GetModuleAddress(evmtypes.ModuleName)
			ugertSender = suite.BankKeeper.GetBalance(suite.Ctx, moduleAddr, "ugert")
			suite.Require().Equal(tt.expModuleCoins.AmountOf("ugert").Int64(), ugertSender.Amount.Int64())
			actualAgert = suite.Keeper.GetBalance(suite.Ctx, moduleAddr)
			suite.Require().Equal(tt.expModuleCoins.AmountOf("agert").Int64(), actualAgert.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestBurnCoins() {
	startingUgert := sdk.NewInt(100)
	tests := []struct {
		name       string
		burnCoins  sdk.Coins
		expUgert   sdk.Int
		expAgert   sdk.Int
		hasErr     bool
		agertStart sdk.Int
	}{
		{
			"burn more than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_021_000_000_002)),
			sdk.NewInt(88),
			sdk.NewInt(100_000_000_000),
			false,
			sdk.NewInt(121_000_000_002),
		},
		{
			"burn less than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 122)),
			sdk.NewInt(100),
			sdk.NewInt(878),
			false,
			sdk.NewInt(1000),
		},
		{
			"burn an exact amount of ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 98_000_000_000_000)),
			sdk.NewInt(2),
			sdk.NewInt(10),
			false,
			sdk.NewInt(10),
		},
		{
			"burn no agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 0)),
			startingUgert,
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"errors if burning other coins",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 500), sdk.NewInt64Coin("busd", 1000)),
			startingUgert,
			sdk.NewInt(100),
			true,
			sdk.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("agert", 12_000_000_000_000),
				sdk.NewInt64Coin("agert", 2_000_000_000_000),
			},
			startingUgert,
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"errors if burn amount is negative",
			sdk.Coins{sdk.Coin{Denom: "agert", Amount: sdk.NewInt(-100)}},
			startingUgert,
			sdk.NewInt(50),
			true,
			sdk.NewInt(50),
		},
		{
			"errors if not enough agert to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100_999_000_000_000)),
			sdk.NewInt(0),
			sdk.NewInt(99_000_000_000),
			true,
			sdk.NewInt(99_000_000_000),
		},
		{
			"errors if not enough ugert to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 200_000_000_000_000)),
			sdk.NewInt(100),
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"converts 1 ugert to agert if not enough agert to cover",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_021_000_000_002)),
			sdk.NewInt(87),
			sdk.NewInt(980_000_000_000),
			false,
			sdk.NewInt(1_000_000_002),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			startingCoins := sdk.NewCoins(
				sdk.NewCoin("ugert", startingUgert),
				sdk.NewCoin("agert", tt.agertStart),
			)
			suite.FundModuleAccountWithgert(evmtypes.ModuleName, startingCoins)

			err := suite.EvmBankKeeper.BurnCoins(suite.Ctx, evmtypes.ModuleName, tt.burnCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check ugert
			ugertActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, "ugert")
			suite.Require().Equal(tt.expUgert, ugertActual.Amount)

			// check agert
			agertActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.expAgert, agertActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestMintCoins() {
	tests := []struct {
		name       string
		mintCoins  sdk.Coins
		ugert      sdk.Int
		agert      sdk.Int
		hasErr     bool
		agertStart sdk.Int
	}{
		{
			"mint more than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_021_000_000_002)),
			sdk.NewInt(12),
			sdk.NewInt(21_000_000_002),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint less than 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 901_000_000_001)),
			sdk.ZeroInt(),
			sdk.NewInt(901_000_000_001),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint an exact amount of ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 123_000_000_000_000_000)),
			sdk.NewInt(123_000),
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint no agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 0)),
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"errors if minting other coins",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.ZeroInt(),
			sdk.NewInt(100),
			true,
			sdk.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("agert", 12_000_000_000_000),
				sdk.NewInt64Coin("agert", 2_000_000_000_000),
			},
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"errors if mint amount is negative",
			sdk.Coins{sdk.Coin{Denom: "agert", Amount: sdk.NewInt(-100)}},
			sdk.ZeroInt(),
			sdk.NewInt(50),
			true,
			sdk.NewInt(50),
		},
		{
			"adds to existing agert balance",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 12_021_000_000_002)),
			sdk.NewInt(12),
			sdk.NewInt(21_000_000_102),
			false,
			sdk.NewInt(100),
		},
		{
			"convert agert balance to ugert if it exceeds 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 10_999_000_000_000)),
			sdk.NewInt(12),
			sdk.NewInt(1_200_000_001),
			false,
			sdk.NewInt(1_002_200_000_001),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundModuleAccountWithgert(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("ugert", 10)))
			suite.FundModuleAccountWithgert(evmtypes.ModuleName, sdk.NewCoins(sdk.NewCoin("agert", tt.agertStart)))

			err := suite.EvmBankKeeper.MintCoins(suite.Ctx, evmtypes.ModuleName, tt.mintCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check ugert
			ugertActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, "ugert")
			suite.Require().Equal(tt.ugert, ugertActual.Amount)

			// check agert
			agertActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.agert, agertActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestValidateEvmCoins() {
	tests := []struct {
		name      string
		coins     sdk.Coins
		shouldErr bool
	}{
		{
			"valid coins",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 500)),
			false,
		},
		{
			"dup coins",
			sdk.Coins{sdk.NewInt64Coin("agert", 500), sdk.NewInt64Coin("agert", 500)},
			true,
		},
		{
			"not evm coins",
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 500)),
			true,
		},
		{
			"negative coins",
			sdk.Coins{sdk.Coin{Denom: "agert", Amount: sdk.NewInt(-500)}},
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := keeper.ValidateEvmCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertOneUgertToAgertIfNeeded() {
	agertNeeded := sdk.NewInt(200)
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
		success       bool
	}{
		{
			"not enough ugert for conversion",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100)),
			false,
		},
		{
			"converts 1 ugert to agert",
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 10), sdk.NewInt64Coin("agert", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 9), sdk.NewInt64Coin("agert", 1_000_000_000_100)),
			true,
		},
		{
			"conversion not needed",
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 10), sdk.NewInt64Coin("agert", 200)),
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 10), sdk.NewInt64Coin("agert", 200)),
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithgert(suite.Addrs[0], tt.startingCoins)
			err := suite.EvmBankKeeper.ConvertOneUgertToAgertIfNeeded(suite.Ctx, suite.Addrs[0], agertNeeded)
			modulegert := suite.BankKeeper.GetBalance(suite.Ctx, suite.AccountKeeper.GetModuleAddress(types.ModuleName), "ugert")
			if tt.success {
				suite.Require().NoError(err)
				if tt.startingCoins.AmountOf("agert").LT(agertNeeded) {
					suite.Require().Equal(sdk.OneInt(), modulegert.Amount)
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(sdk.ZeroInt(), modulegert.Amount)
			}

			agert := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf("agert"), agert)
			ugert := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ugert")
			suite.Require().Equal(tt.expectedCoins.AmountOf("ugert"), ugert.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertAgertToUgert() {
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
	}{
		{
			"not enough ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 100), sdk.NewInt64Coin("ugert", 0)),
		},
		{
			"converts agert for 1 ugert",
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 10), sdk.NewInt64Coin("agert", 1_000_000_000_003)),
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 11), sdk.NewInt64Coin("agert", 3)),
		},
		{
			"converts more than 1 ugert of agert",
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 10), sdk.NewInt64Coin("agert", 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 18), sdk.NewInt64Coin("agert", 123)),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			err := suite.App.FundModuleAccount(suite.Ctx, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("ugert", 10)))
			suite.Require().NoError(err)
			suite.FundAccountWithgert(suite.Addrs[0], tt.startingCoins)
			err = suite.EvmBankKeeper.ConvertAgertToUgert(suite.Ctx, suite.Addrs[0])
			suite.Require().NoError(err)
			agert := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf("agert"), agert)
			ugert := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "ugert")
			suite.Require().Equal(tt.expectedCoins.AmountOf("ugert"), ugert.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSplitAgertCoins() {
	tests := []struct {
		name          string
		coins         sdk.Coins
		expectedCoins sdk.Coins
		shouldErr     bool
	}{
		{
			"invalid coins",
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 500)),
			nil,
			true,
		},
		{
			"empty coins",
			sdk.NewCoins(),
			sdk.NewCoins(),
			false,
		},
		{
			"ugert & agert coins",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 8), sdk.NewInt64Coin("agert", 123)),
			false,
		},
		{
			"only agert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 10_123)),
			sdk.NewCoins(sdk.NewInt64Coin("agert", 10_123)),
			false,
		},
		{
			"only ugert",
			sdk.NewCoins(sdk.NewInt64Coin("agert", 5_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("ugert", 5)),
			false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ugert, agert, err := keeper.SplitAgertCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tt.expectedCoins.AmountOf("ugert"), ugert.Amount)
				suite.Require().Equal(tt.expectedCoins.AmountOf("agert"), agert)
			}
		})
	}
}

func TestEvmBankKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(evmBankKeeperTestSuite))
}

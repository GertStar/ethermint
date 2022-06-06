package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/Gert-Chain/core/app"
	"github.com/Gert-Chain/core/x/evmutil/keeper"
)

type Suite struct {
	suite.Suite

	App           app.TestApp
	Ctx           sdk.Context
	BankKeeper    bankkeeper.Keeper
	AccountKeeper authkeeper.AccountKeeper
	Keeper        keeper.Keeper
	EvmBankKeeper keeper.EvmBankKeeper
	Addrs         []sdk.AccAddress
	EvmModuleAddr sdk.AccAddress
}

func (suite *Suite) SetupTest() {
	tApp := app.NewTestApp()

	suite.Ctx = tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.App = tApp
	suite.BankKeeper = tApp.GetBankKeeper()
	suite.AccountKeeper = tApp.GetAccountKeeper()
	suite.Keeper = tApp.GetEvmutilKeeper()
	suite.EvmBankKeeper = keeper.NewEvmBankKeeper(tApp.GetEvmutilKeeper(), suite.BankKeeper, suite.AccountKeeper)
	suite.EvmModuleAddr = suite.AccountKeeper.GetModuleAddress(evmtypes.ModuleName)

	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	suite.Addrs = addrs

	evmGenesis := evmtypes.DefaultGenesisState()
	evmGenesis.Params.EvmDenom = "agert"
	gs := app.GenesisState{evmtypes.ModuleName: suite.App.AppCodec().MustMarshalJSON(evmGenesis)}
	suite.App.InitializeFromGenesisStates(gs)
}

func (suite *Suite) FundAccountWithgert(addr sdk.AccAddress, coins sdk.Coins) {
	ugert := coins.AmountOf("ugert")
	if ugert.IsPositive() {
		err := suite.App.FundAccount(suite.Ctx, addr, sdk.NewCoins(sdk.NewCoin("ugert", ugert)))
		suite.Require().NoError(err)
	}
	agert := coins.AmountOf("agert")
	if agert.IsPositive() {
		err := suite.Keeper.SetBalance(suite.Ctx, addr, agert)
		suite.Require().NoError(err)
	}
}

func (suite *Suite) FundModuleAccountWithgert(moduleName string, coins sdk.Coins) {
	ugert := coins.AmountOf("ugert")
	if ugert.IsPositive() {
		err := suite.App.FundModuleAccount(suite.Ctx, moduleName, sdk.NewCoins(sdk.NewCoin("ugert", ugert)))
		suite.Require().NoError(err)
	}
	agert := coins.AmountOf("agert")
	if agert.IsPositive() {
		addr := suite.AccountKeeper.GetModuleAddress(moduleName)
		err := suite.Keeper.SetBalance(suite.Ctx, addr, agert)
		suite.Require().NoError(err)
	}
}

package ante_test

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/Gert-Chain/core/app"
	"github.com/Gert-Chain/core/app/ante"
)

func mustParseDecCoins(value string) sdk.DecCoins {
	coins, err := sdk.ParseDecCoins(strings.ReplaceAll(value, ";", ","))
	if err != nil {
		panic(err)
	}

	return coins
}

func TestEvmMinGasFilter(t *testing.T) {
	tApp := app.NewTestApp()
	handler := ante.NewEvmMinGasFilter(tApp.GetEvmKeeper())

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.GetEvmKeeper().SetParams(ctx, evmtypes.Params{
		EvmDenom: "agert",
	})

	testCases := []struct {
		name                 string
		minGasPrices         sdk.DecCoins
		expectedMinGasPrices sdk.DecCoins
	}{
		{
			"no min gas prices",
			mustParseDecCoins(""),
			mustParseDecCoins(""),
		},
		{
			"zero ugert gas price",
			mustParseDecCoins("0ugert"),
			mustParseDecCoins("0ugert"),
		},
		{
			"non-zero ugert gas price",
			mustParseDecCoins("0.001ugert"),
			mustParseDecCoins("0.001ugert"),
		},
		{
			"zero ugert gas price, min agert price",
			mustParseDecCoins("0ugert;100000agert"),
			mustParseDecCoins("0ugert"), // agert is removed
		},
		{
			"zero ugert gas price, min agert price, other token",
			mustParseDecCoins("0ugert;100000agert;0.001other"),
			mustParseDecCoins("0ugert;0.001other"), // agert is removed
		},
		{
			"non-zero ugert gas price, min agert price",
			mustParseDecCoins("0.25ugert;100000agert;0.001other"),
			mustParseDecCoins("0.25ugert;0.001other"), // agert is removed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

			ctx = ctx.WithMinGasPrices(tc.minGasPrices)
			mmd := MockAnteHandler{}

			_, err := handler.AnteHandle(ctx, nil, false, mmd.AnteHandle)
			require.NoError(t, err)
			require.True(t, mmd.WasCalled)

			assert.NoError(t, mmd.CalledCtx.MinGasPrices().Validate())
			assert.Equal(t, tc.expectedMinGasPrices, mmd.CalledCtx.MinGasPrices())
		})
	}
}

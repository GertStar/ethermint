package types_test

import (
	"github.com/Gert-Chain/core/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func init() {
	gertConfig := sdk.GetConfig()
	app.SetBech32AddressPrefixes(gertConfig)
	app.SetBip44CoinType(gertConfig)
	gertConfig.Seal()
}

package gertdist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Gert-Chain/core/x/gertdist/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}

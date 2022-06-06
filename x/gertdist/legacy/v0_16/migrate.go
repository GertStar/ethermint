package v0_16

import (
	v015gertdist "github.com/Gert-Chain/core/x/gertdist/legacy/v0_15"
	v016gertdist "github.com/Gert-Chain/core/x/gertdist/types"
)

func migrateParams(oldParams v015gertdist.Params) v016gertdist.Params {
	periods := make([]v016gertdist.Period, len(oldParams.Periods))
	for i, oldPeriod := range oldParams.Periods {
		periods[i] = v016gertdist.Period{
			Start:     oldPeriod.Start,
			End:       oldPeriod.End,
			Inflation: oldPeriod.Inflation,
		}
	}
	return v016gertdist.Params{
		Periods: periods,
		Active:  oldParams.Active,
	}
}

// Migrate converts v0.15 gertdist state and returns it in v0.16 format
func Migrate(oldState v015gertdist.GenesisState) *v016gertdist.GenesisState {
	return &v016gertdist.GenesisState{
		Params:            migrateParams(oldState.Params),
		PreviousBlockTime: oldState.PreviousBlockTime,
	}
}

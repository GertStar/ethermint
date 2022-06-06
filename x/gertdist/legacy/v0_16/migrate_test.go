package v0_16

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmtime "github.com/tendermint/tendermint/types/time"

	app "github.com/Gert-Chain/core/app"
	v015gertdist "github.com/Gert-Chain/core/x/gertdist/legacy/v0_15"
	v016gertdist "github.com/Gert-Chain/core/x/gertdist/types"
)

type migrateTestSuite struct {
	suite.Suite

	addresses   []sdk.AccAddress
	v15genstate v015gertdist.GenesisState
	cdc         codec.Codec
	legacyCdc   *codec.LegacyAmino
}

func (s *migrateTestSuite) SetupTest() {
	app.SetSDKConfig()

	s.v15genstate = v015gertdist.GenesisState{
		Params:            v015gertdist.Params{},
		PreviousBlockTime: tmtime.Canonical(time.Unix(1, 0)),
	}

	config := app.MakeEncodingConfig()
	s.cdc = config.Marshaler

	legacyCodec := codec.NewLegacyAmino()
	s.legacyCdc = legacyCodec

	_, accAddresses := app.GeneratePrivKeyAddressPairs(10)
	s.addresses = accAddresses
}

func (s *migrateTestSuite) TestMigrate_JSON() {
	// Migrate v15 gertdist to v16
	data := `{
		"params": {
			"active": true,
			"periods": [
				{
					"end": "2021-06-01T14:00:00Z",
					"inflation": "1.000000000000000000",
					"start": "2020-06-01T14:00:00Z"
				},
				{
					"end": "2022-06-01T14:00:00Z",
					"inflation": "1.000000002293273137",
					"start": "2021-06-01T14:00:00Z"
				},
				{
					"end": "2023-06-01T14:00:00Z",
					"inflation": "1.000000001167363430",
					"start": "2022-06-01T14:00:00Z"
				},
				{
					"end": "2024-06-01T14:00:00Z",
					"inflation": "1.000000000782997609",
					"start": "2023-06-01T14:00:00Z"
				}
			]
		},
		"previous_block_time": "2021-11-05T21:13:12.85608847Z"
	}`
	err := s.legacyCdc.UnmarshalJSON([]byte(data), &s.v15genstate)
	s.Require().NoError(err)
	genstate := Migrate(s.v15genstate)

	// Compare expect v16 gertdist json with migrated json
	actual := s.cdc.MustMarshalJSON(genstate)
	expected := `{
		"params": {
			"active": true,
			"periods": [
				{
					"start": "2020-06-01T14:00:00Z",
					"end": "2021-06-01T14:00:00Z",
					"inflation": "1.000000000000000000"
				},
				{
					"start": "2021-06-01T14:00:00Z",
					"end": "2022-06-01T14:00:00Z",
					"inflation": "1.000000002293273137"
				},
				{
					"start": "2022-06-01T14:00:00Z",
					"end": "2023-06-01T14:00:00Z",
					"inflation": "1.000000001167363430"
				},
				{
					"start": "2023-06-01T14:00:00Z",
					"end": "2024-06-01T14:00:00Z",
					"inflation": "1.000000000782997609"
				}
			]
		},
		"previous_block_time": "2021-11-05T21:13:12.856088470Z"
	}`
	s.Require().NoError(err)
	s.Require().JSONEq(string(expected), string(actual))
}

func (s *migrateTestSuite) TestMigrate_Params() {
	params := v015gertdist.Params{
		Active: true,
		Periods: v015gertdist.Periods{
			{
				Start:     time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("0.1"),
			},
			{
				Start:     time.Date(2010, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2011, time.March, 2, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("0.01"),
			},
		},
	}
	expected := v016gertdist.Params{
		Active: true,
		Periods: []v016gertdist.Period{
			{
				Start:     time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("0.1"),
			},
			{
				Start:     time.Date(2010, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2011, time.March, 2, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("0.01"),
			},
		},
	}

	s.v15genstate.Params = params
	genState := Migrate(s.v15genstate)
	s.Require().Equal(expected, genState.Params)
}

func TestMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(migrateTestSuite))
}

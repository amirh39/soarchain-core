package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"soarchain/x/poa/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ClientList: []types.Client{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ChallengerList: []types.Challenger{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				RunnerList: []types.Runner{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				GuardList: []types.Guard{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated client",
			genState: &types.GenesisState{
				ClientList: []types.Client{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated challenger",
			genState: &types.GenesisState{
				ChallengerList: []types.Challenger{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated runner",
			genState: &types.GenesisState{
				RunnerList: []types.Runner{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated guard",
			genState: &types.GenesisState{
				GuardList: []types.Guard{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

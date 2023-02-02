package types_test

import (
	"testing"

	"soarchain/x/poa/types"

	"github.com/stretchr/testify/require"
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

				VrfDataList: []types.VrfData{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				VrfUserList: []types.VrfUser{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				EpochData: types.EpochData{
					TotalEpochs: 81,
					EpochV2VRX:  "46",
					EpochV2VBX:  "26",
					EpochV2NBX:  "99",
					EpochRunner: "26",
				},
				MotusWalletList: []types.MotusWallet{
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

		{
			desc: "duplicated vrfData",
			genState: &types.GenesisState{
				VrfDataList: []types.VrfData{
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
			desc: "duplicated vrfUser",
			genState: &types.GenesisState{
				VrfUserList: []types.VrfUser{
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
			desc: "duplicated motusWallet",
			genState: &types.GenesisState{
				MotusWalletList: []types.MotusWallet{
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

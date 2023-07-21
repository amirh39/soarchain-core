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
						PubKey: "0",
					},
					{
						PubKey: "1",
					},
				},
				RunnerList: []types.Runner{
					{
						PubKey: "0",
					},
					{
						PubKey: "1",
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
				MotusWalletList: []types.MotusWallet{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				MasterKey: types.MasterKey{
					MasterCertificate: "26",
					MasterAccount:     "70",
				},
				FactoryKeysList: []types.FactoryKeys{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				FactoryKeysCount: 2,
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
						PubKey: "0",
					},
					{
						PubKey: "0",
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
						PubKey: "0",
					},
					{
						PubKey: "0",
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
		{
			desc: "duplicated factoryKeys",
			genState: &types.GenesisState{
				FactoryKeysList: []types.FactoryKeys{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid factoryKeys count",
			genState: &types.GenesisState{
				FactoryKeysList: []types.FactoryKeys{
					{
						Id: 1,
					},
				},
				FactoryKeysCount: 0,
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

package types_test

import (
	"testing"

	"github.com/soar-robotics/soarchain-core/x/poa/types"

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
				VrfDataList: []types.VrfData{
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

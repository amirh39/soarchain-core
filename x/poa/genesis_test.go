package poa_test

import (
	"testing"

	keepertest "github.com/soar-robotics/soarchain-core/testutil/keeper"
	"github.com/soar-robotics/soarchain-core/testutil/nullify"
	"github.com/soar-robotics/soarchain-core/x/poa"
	"github.com/soar-robotics/soarchain-core/x/poa/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ReputationList: []types.Reputation{
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
		MasterKey: types.MasterKey{
			MasterCertificate: "39",
			MasterAccount:     "62",
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PoaKeeper(t)
	poa.InitGenesis(ctx, *k, genesisState)
	got := poa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ReputationList, got.ReputationList)
	require.ElementsMatch(t, genesisState.VrfDataList, got.VrfDataList)
	require.Equal(t, genesisState.MasterKey, got.MasterKey)
	require.ElementsMatch(t, genesisState.FactoryKeysList, got.FactoryKeysList)
	require.Equal(t, genesisState.FactoryKeysCount, got.FactoryKeysCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

package poa_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa"
	"soarchain/x/poa/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		TotalClients: types.TotalClients{
			Count: 47,
		},
		TotalChallengers: types.TotalChallengers{
			Count: 59,
		},
		TotalRunners: types.TotalRunners{
			Count: 78,
		},
		ChallengerByIndexList: []types.ChallengerByIndex{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PoaKeeper(t)
	poa.InitGenesis(ctx, *k, genesisState)
	got := poa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClientList, got.ClientList)
	require.ElementsMatch(t, genesisState.ChallengerList, got.ChallengerList)
	require.ElementsMatch(t, genesisState.RunnerList, got.RunnerList)
	require.ElementsMatch(t, genesisState.GuardList, got.GuardList)
	require.Equal(t, genesisState.TotalClients, got.TotalClients)
	require.Equal(t, genesisState.TotalChallengers, got.TotalChallengers)
	require.Equal(t, genesisState.TotalRunners, got.TotalRunners)
	require.ElementsMatch(t, genesisState.ChallengerByIndexList, got.ChallengerByIndexList)
	// this line is used by starport scaffolding # genesis/test/assert
}

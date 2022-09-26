package poa_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa"
	"soarchain/x/poa/types"
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
	// this line is used by starport scaffolding # genesis/test/assert
}

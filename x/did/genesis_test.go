package did_test

import (
	"soarchain/testutil/nullify"
	"soarchain/x/did"
	"soarchain/x/did/types"
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.DidKeeper(t)
	did.InitGenesis(ctx, *k, genesisState)
	got := did.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClientDidList, got.ClientDidList)
	require.ElementsMatch(t, genesisState.RunnerDidList, got.RunnerDidList)
	require.ElementsMatch(t, genesisState.ChallengerDidList, got.ChallengerDidList)
}

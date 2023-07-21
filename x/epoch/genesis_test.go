package epoch_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/epoch"
	"soarchain/x/epoch/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.EpochKeeper(t)
	epoch.InitGenesis(ctx, *k, genesisState)
	got := epoch.ExportGenesis(ctx, *k)
	require.NotNil(t, got)
	require.Equal(t, genesisState.EpochData, got.EpochData)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

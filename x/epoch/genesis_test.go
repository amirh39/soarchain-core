package epoch_test

import (
	"testing"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/testutil/nullify"
	"github.com/amirh39/soarchain-core/x/epoch"
	"github.com/amirh39/soarchain-core/x/epoch/types"

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

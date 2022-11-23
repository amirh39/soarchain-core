package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

func createTestTotalRunners(keeper *keeper.Keeper, ctx sdk.Context) types.TotalRunners {
	item := types.TotalRunners{}
	keeper.SetTotalRunners(ctx, item)
	return item
}

func TestTotalRunnersGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createTestTotalRunners(keeper, ctx)
	rst, found := keeper.GetTotalRunners(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTotalRunnersRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	createTestTotalRunners(keeper, ctx)
	keeper.RemoveTotalRunners(ctx)
	_, found := keeper.GetTotalRunners(ctx)
	require.False(t, found)
}

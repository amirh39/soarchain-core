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

func createTestEpochData(keeper *keeper.Keeper, ctx sdk.Context) types.EpochData {
	item := types.EpochData{}
	keeper.SetEpochData(ctx, item)
	return item
}

func TestEpochDataGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createTestEpochData(keeper, ctx)
	rst, found := keeper.GetEpochData(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestEpochDataRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	createTestEpochData(keeper, ctx)
	keeper.RemoveEpochData(ctx)
	_, found := keeper.GetEpochData(ctx)
	require.False(t, found)
}

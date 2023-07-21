package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
)

func Test_EpochDataGet(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	item := CreateEpochData(keeper, ctx)
	rst, found := keeper.GetEpochData(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func Test_EpochDataRemove(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	CreateEpochData(keeper, ctx)
	keeper.RemoveEpochData(ctx)
	_, found := keeper.GetEpochData(ctx)
	require.False(t, found)
}

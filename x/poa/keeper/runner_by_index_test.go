package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRunnerByIndex(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RunnerByIndex {
	items := make([]types.RunnerByIndex, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRunnerByIndex(ctx, items[i])
	}
	return items
}

func TestRunnerByIndexGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunnerByIndex(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRunnerByIndex(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRunnerByIndexRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunnerByIndex(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRunnerByIndex(ctx,
			item.Index,
		)
		_, found := keeper.GetRunnerByIndex(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRunnerByIndexGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunnerByIndex(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRunnerByIndex(ctx)),
	)
}

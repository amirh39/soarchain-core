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

func createNRunner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRunner(ctx, items[i])
	}
	return items
}

func TestRunnerGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRunner(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRunnerRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRunner(ctx,
			item.Index,
		)
		_, found := keeper.GetRunner(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRunnerGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRunner(ctx)),
	)
}

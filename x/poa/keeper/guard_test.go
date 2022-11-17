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

func createNGuard(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Guard {
	items := make([]types.Guard, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetGuard(ctx, items[i])
	}
	return items
}

func TestGuardGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNGuard(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGuard(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGuardRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNGuard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGuard(ctx,
			item.Index,
		)
		_, found := keeper.GetGuard(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGuardGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNGuard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGuard(ctx)),
	)
}

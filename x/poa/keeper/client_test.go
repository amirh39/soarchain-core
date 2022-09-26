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

func createNClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClient(ctx, items[i])
	}
	return items
}

func TestClientGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNClient(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClient(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClientRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNClient(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClient(ctx,
			item.Index,
		)
		_, found := keeper.GetClient(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestClientGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNClient(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClient(ctx)),
	)
}

package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func Test_ClientGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNClient(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClient(ctx,
			item.PubKey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func Test_ClientRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNClient(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClient(ctx,
			item.PubKey,
		)
		_, found := keeper.GetClient(ctx,
			item.PubKey,
		)
		require.False(t, found)
	}
}

func Test_ClientGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNClient(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClient(ctx)),
	)
}

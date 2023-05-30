package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func TestVrfDataGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNVrfData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVrfData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVrfDataRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNVrfData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVrfData(ctx,
			item.Index,
		)
		_, found := keeper.GetVrfData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestVrfDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNVrfData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVrfData(ctx)),
	)
}

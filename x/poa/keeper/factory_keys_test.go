package keeper_test

import (
	"testing"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func Test_FactoryKeysGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNFactoryKeys(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFactoryKeys(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func Test_FactoryKeysRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNFactoryKeys(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFactoryKeys(ctx, item.Id)
		_, found := keeper.GetFactoryKeys(ctx, item.Id)
		require.False(t, found)
	}
}

func Test_FactoryKeysGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNFactoryKeys(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFactoryKeys(ctx)),
	)
}

func Test_FactoryKeysCount(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNFactoryKeys(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFactoryKeysCount(ctx))
}

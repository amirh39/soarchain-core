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

func createNFactoryKeys(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FactoryKeys {
	items := make([]types.FactoryKeys, n)
	for i := range items {
		items[i].Id = keeper.AppendFactoryKeys(ctx, items[i])
	}
	return items
}

func TestFactoryKeysGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNFactoryKeys(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFactoryKeys(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFactoryKeysRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNFactoryKeys(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFactoryKeys(ctx, item.Id)
		_, found := keeper.GetFactoryKeys(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFactoryKeysGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNFactoryKeys(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFactoryKeys(ctx)),
	)
}

func TestFactoryKeysCount(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNFactoryKeys(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFactoryKeysCount(ctx))
}

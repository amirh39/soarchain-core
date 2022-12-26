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

func createNVrfData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VrfData {
	items := make([]types.VrfData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetVrfData(ctx, items[i])
	}
	return items
}

func TestVrfDataGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNVrfData(keeper, ctx, 10)
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
	items := createNVrfData(keeper, ctx, 10)
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
	items := createNVrfData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVrfData(ctx)),
	)
}

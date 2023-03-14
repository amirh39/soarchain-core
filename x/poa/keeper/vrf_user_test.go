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

func createNVrfUser(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VrfUser {
	items := make([]types.VrfUser, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetVrfUser(ctx, items[i])
	}
	return items
}

func TestVrfUserGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNVrfUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVrfUser(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVrfUserRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNVrfUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVrfUser(ctx,
			item.Index,
		)
		_, found := keeper.GetVrfUser(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestVrfUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNVrfUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVrfUser(ctx)),
	)
}

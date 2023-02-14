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

func createNMotusWallet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MotusWallet {
	items := make([]types.MotusWallet, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMotusWallet(ctx, items[i])
	}
	return items
}

func TestMotusWalletGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNMotusWallet(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMotusWallet(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMotusWalletRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNMotusWallet(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMotusWallet(ctx,
			item.Index,
		)
		_, found := keeper.GetMotusWallet(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestMotusWalletGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNMotusWallet(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMotusWallet(ctx)),
	)
}

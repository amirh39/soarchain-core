package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func Test_MotusWalletGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNMotusWallet(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMotusWallet(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func Test_MotusWalletRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNMotusWallet(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMotusWallet(ctx,
			item.Address,
		)
		_, found := keeper.GetMotusWallet(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func Test_MotusWalletGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNMotusWallet(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMotusWallet(ctx)),
	)
}

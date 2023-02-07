package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/soarmint/keeper"
	"soarchain/x/soarmint/types"
)

func createTestMinter(keeper *keeper.Keeper, ctx sdk.Context) types.Minter {
	item := types.Minter{}
	keeper.SetMinter(ctx, item)
	return item
}

func TestMinterGet(t *testing.T) {
	keeper, ctx := keepertest.SoarmintKeeper(t)
	item := createTestMinter(keeper, ctx)
	rst, found := keeper.GetMinter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMinterRemove(t *testing.T) {
	keeper, ctx := keepertest.SoarmintKeeper(t)
	createTestMinter(keeper, ctx)
	keeper.RemoveMinter(ctx)
	_, found := keeper.GetMinter(ctx)
	require.False(t, found)
}

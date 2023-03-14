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

func createTestMasterKey(keeper *keeper.Keeper, ctx sdk.Context) types.MasterKey {
	item := types.MasterKey{}
	keeper.SetMasterKey(ctx, item)
	return item
}

func TestMasterKeyGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createTestMasterKey(keeper, ctx)
	rst, found := keeper.GetMasterKey(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMasterKeyRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	createTestMasterKey(keeper, ctx)
	keeper.RemoveMasterKey(ctx)
	_, found := keeper.GetMasterKey(ctx)
	require.False(t, found)
}

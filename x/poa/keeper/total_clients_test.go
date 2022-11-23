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

func createTestTotalClients(keeper *keeper.Keeper, ctx sdk.Context) types.TotalClients {
	item := types.TotalClients{}
	keeper.SetTotalClients(ctx, item)
	return item
}

func TestTotalClientsGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createTestTotalClients(keeper, ctx)
	rst, found := keeper.GetTotalClients(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTotalClientsRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	createTestTotalClients(keeper, ctx)
	keeper.RemoveTotalClients(ctx)
	_, found := keeper.GetTotalClients(ctx)
	require.False(t, found)
}

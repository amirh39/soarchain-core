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

func createTestTotalChallengers(keeper *keeper.Keeper, ctx sdk.Context) types.TotalChallengers {
	item := types.TotalChallengers{}
	keeper.SetTotalChallengers(ctx, item)
	return item
}

func TestTotalChallengersGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := createTestTotalChallengers(keeper, ctx)
	rst, found := keeper.GetTotalChallengers(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTotalChallengersRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	createTestTotalChallengers(keeper, ctx)
	keeper.RemoveTotalChallengers(ctx)
	_, found := keeper.GetTotalChallengers(ctx)
	require.False(t, found)
}

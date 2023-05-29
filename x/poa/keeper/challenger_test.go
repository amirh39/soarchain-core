package keeper_test

import (
	"strconv"
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNChallenger(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenger {
	items := make([]types.Challenger, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetChallenger(ctx, items[i])
	}
	return items
}

func createV2NTypeChallenger(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenger {
	items := make([]types.Challenger, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Type = "v2n"

		keeper.SetChallenger(ctx, items[i])
	}
	return items
}

func TestChallengerGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallenger(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChallenger(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChallengerRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallenger(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChallenger(ctx,
			item.Address,
		)
		_, found := keeper.GetChallenger(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestChallengerGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallenger(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChallenger(ctx)),
	)
}

func TestGetChallengerUsingPubKey(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	challengers := createNChallenger(keeper, ctx, 10)

	// Pick a random challenger and set its PubKey
	targetChallenger := challengers[4]
	targetChallenger.PubKey = "test-pubkey"
	keeper.SetChallenger(ctx, targetChallenger)

	// Test that GetChallengerUsingPubKey returns the correct challenger
	result, _ := keeper.GetChallengerUsingPubKey(ctx, targetChallenger.PubKey)
	require.Equal(t, targetChallenger, result)
}

func Test_GetChallengerByType(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createV2NTypeChallenger(keeper, ctx, 1)
	for _, item := range items {
		rst, found := keeper.GetChallengerByType(ctx,
			item.Address,
			item.Type,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

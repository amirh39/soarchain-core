package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func Test_ChallengerGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNChallenger(keeper, ctx, 10)
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
func Test_ChallengerRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNChallenger(keeper, ctx, 10)
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

func Test_ChallengerGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNChallenger(keeper, ctx, 10)
	require.Equal(t,
		len(items),
		10,
	)
}

func Test_GetChallengerUsingPubKey(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	challengers := CreateNChallenger(keeper, ctx, 10)

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
	items := CreateV2NTypeChallenger(keeper, ctx, 1)
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

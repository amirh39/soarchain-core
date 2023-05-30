package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func TestRunnerGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNRunner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRunner(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRunnerRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNRunner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRunner(ctx,
			item.Address,
		)
		_, found := keeper.GetRunner(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestRunnerGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNRunner(keeper, ctx, 10)
	require.Equal(t,
		len(items),
		10,
	)
}

func TestGetRunnerUsingPubKey(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := CreateNRunner(keeper, ctx, 10)

	// Pick a random challenger and set its PubKey
	targetRunnner := runners[4]
	targetRunnner.PubKey = "test-pubkey"
	keeper.SetRunner(ctx, targetRunnner)

	// Test that GetChallengerUsingPubKey returns the correct challenger
	result, _ := keeper.GetRunnerUsingPubKey(ctx, targetRunnner.PubKey)
	require.Equal(t, targetRunnner, result)
}

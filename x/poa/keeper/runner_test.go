package keeper_test

import (
	"strconv"
	"testing"
	"time"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRunner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetRunner(ctx, items[i])
	}
	return items
}

func createAChallengableRunner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Score = "80"
		items[i].LastTimeChallenged = "2022-01-06 11:05:17.40125 +0000 UTC"
		items[i].CoolDownTolerance = "1"

		keeper.SetRunner(ctx, items[i])
	}
	return items
}

func TestRunnerGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNRunner(keeper, ctx, 10)
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
	items := createNRunner(keeper, ctx, 10)
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
	items := createNRunner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRunner(ctx)),
	)
}

func TestGetRunnerUsingPubKey(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := createNRunner(keeper, ctx, 10)

	// Pick a random challenger and set its PubKey
	targetRunnner := runners[4]
	targetRunnner.PubKey = "test-pubkey"
	keeper.SetRunner(ctx, targetRunnner)

	// Test that GetChallengerUsingPubKey returns the correct challenger
	result, _ := keeper.GetRunnerUsingPubKey(ctx, targetRunnner.PubKey)
	require.Equal(t, targetRunnner, result)
}

func Test_GetChallengeableRunner(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	ctx = ctx.WithBlockTime(time.Date(2023, 01, 06, 11, 05, 17, 40125, time.UTC))
	items := createAChallengableRunner(keeper, ctx, 1)
	for _, item := range items {
		rst, found := keeper.GetChallengeableRunner(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

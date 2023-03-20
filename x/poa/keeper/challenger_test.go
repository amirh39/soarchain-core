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
		items[i].PubKey = strconv.Itoa(i)

		keeper.SetChallenger(ctx, items[i])
	}
	return items
}

func TestChallengerGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallenger(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChallenger(ctx,
			item.PubKey,
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
			item.PubKey,
		)
		_, found := keeper.GetChallenger(ctx,
			item.PubKey,
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

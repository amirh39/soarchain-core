package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNChallengerByIndex(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChallengerByIndex {
	items := make([]types.ChallengerByIndex, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetChallengerByIndex(ctx, items[i])
	}
	return items
}

func TestChallengerByIndexGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallengerByIndex(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChallengerByIndex(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChallengerByIndexRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallengerByIndex(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChallengerByIndex(ctx,
			item.Index,
		)
		_, found := keeper.GetChallengerByIndex(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestChallengerByIndexGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := createNChallengerByIndex(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChallengerByIndex(ctx)),
	)
}

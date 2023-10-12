package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"

	"github.com/stretchr/testify/require"
)

func Test_ReputationGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNReputation(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetReputation(ctx,
			item.PubKey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func Test_ReputationGetAll(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNReputation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllReputation(ctx)),
	)
}

func Test_ReputationByAddressGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	items := CreateNReputation(keeper, ctx, 1)
	for _, item := range items {
		response, found := keeper.GetReputation(ctx,
			item.Address,
		)
		t.Log("response------>", response)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&response),
		)
	}
}

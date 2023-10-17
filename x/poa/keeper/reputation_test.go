package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"

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

func TestGetReputationByType(t *testing.T) {
	k, ctx := keepertest.PoaKeeper(t)

	// Create a test reputation object
	testReputation := types.Reputation{
		PubKey:  "pubkey",
		Address: "testAddress",
		Type:    "testType",
	}

	// Store the test reputation object
	k.SetReputation(ctx, testReputation)

	// Test case 1: Correct reputation type and address
	val, found := k.GetReputationsByAddressAndType(ctx, "testAddress", "testType")
	require.True(t, found)
	require.Equal(t, testReputation, val)

	// Test case 2: Incorrect reputation type
	val, found = k.GetReputationsByAddressAndType(ctx, "testAddress", "wrongType")
	require.False(t, found)
	require.Equal(t, types.Reputation{}, val)

	// Test case 3: Incorrect address
	val, found = k.GetReputationsByAddressAndType(ctx, "wrongAddress", "testType")
	require.False(t, found)
	require.Equal(t, types.Reputation{}, val)
}

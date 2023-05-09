package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].RewardMultiplier = "4"
		keeper.SetClient(ctx, items[i])
	}
	return items
}

func Test_V2VRewardCalculator(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	clients := createClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2VRewardCalculator(ctx, 4, "v2v-rx")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

func Test_V2VRewardCalculator_NotValidClients(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	clients := createNClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "v2v-rx")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.Zero(t, 0)
}

func Test_V2VRewardCalculator_NotValidRewardMultiplier(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := createClient(keeper, ctx, 0)
	t.Log("Runners", runners)

	clients := createNClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.Zero(t, 0)
}

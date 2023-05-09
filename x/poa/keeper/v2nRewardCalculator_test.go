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

func createNRunners(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].IpAddr = "45.12.65.78"
		items[i].RewardMultiplier = "4"

		keeper.SetRunner(ctx, items[i])
	}
	return items
}

func Test_V2NRewardCalculator(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := createNRunners(keeper, ctx, 12)
	t.Log("Runners", runners)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

func Test_V2NRewardCalculator_NotValidRunners(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := createNRunner(keeper, ctx, 0)
	t.Log("Runners", runners)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.Zero(t, 0)
}

func Test_V2NRewardCalculator_NotValidClients(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := createNRunners(keeper, ctx, 0)
	t.Log("Runners", runners)

	clients := createNClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.Zero(t, 0)

}

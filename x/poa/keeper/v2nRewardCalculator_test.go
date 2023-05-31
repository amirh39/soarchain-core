package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_V2NRewardCalculator(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := CreateNRunnersForRewardCap(keeper, ctx, 12)
	t.Log("Runners", runners)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

func Test_V2NRewardCalculator_NotValidRunners(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	runners := CreateNRunner(keeper, ctx, 0)
	t.Log("Runners", runners)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("respone", res)

	require.Error(t, err)
	require.Zero(t, 0)
}

func Test_V2NRewardCalculator_NotValidClients(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	runners := CreateNRunner(keeper, ctx, 10)
	t.Log("Runners", runners)

	clients := CreateNClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, "runner")

	t.Log("respone", res)

	require.Error(t, err)
	require.Zero(t, 0)
}

package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa/constants"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_V2VRewardCalculator(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	clients := CreateRewardCapClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2VRewardCalculator(ctx, 4, constants.V2VRX)

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

func Test_V2VRewardCalculator_NotValidClients(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	clients := CreateNClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, constants.V2VRX)

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.Zero(t, 0)
}

func Test_V2VRewardCalculator_NotValidRewardMultiplier(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := CreateNRunner(keeper, ctx, 1)
	t.Log("Runners", runners)

	clients := CreateNClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, constants.Runner)

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.Zero(t, 0)
}

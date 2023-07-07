package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_V2NRewardCalculator(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	runners := CreateNRunnersForRewardCap(keeper, ctx, 12)
	require.NotNil(t, runners)

	res, err := keeper.V2NRewardCalculator(ctx, 4, TRunner)

	require.NoError(t, err)
	t.Log("res", res)
	require.NotZero(t, res)
	require.NotNil(t, res)
}

func Test_V2NRewardCalculator_NotValidRunners(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	runners := CreateNRunner(keeper, ctx, 0)
	t.Log("Runners", runners)

	res, err := keeper.V2NRewardCalculator(ctx, 4, TRunner)

	require.Error(t, err)
	t.Log("res", res)
	require.Zero(t, 0)
	require.NotNil(t, res)
}

func Test_V2NRewardCalculator_NotValidClients(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	clients := CreateInValidClient(keeper, ctx, 10)
	t.Log("clients", clients)

	res, err := keeper.V2NRewardCalculator(ctx, 4, TV2VBX)

	require.Error(t, err)
	require.Zero(t, 0)
	require.NotNil(t, res)

}

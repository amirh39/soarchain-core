package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
)

func Test_GenerateRandomNumberGet(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	randomNumber, found := keeper.RandomNumber(ctx, 12)
	t.Log("random number ---> ", randomNumber)
	require.True(t, found)
	require.NotNil(t, randomNumber)
}

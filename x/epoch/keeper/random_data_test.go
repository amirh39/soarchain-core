package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
)

func Test_RandomDataGet(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	item := CreateRandomData(keeper, ctx)
	keeper.SetRandomData(ctx, item)
	t.Log("item --->", item)
	require.NotNil(t, item)

	rst, found := keeper.GetRandomData(ctx, "353545345")
	t.Log("random data ---> ", rst)
	require.True(t, found)
}

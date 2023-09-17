package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DeactivateDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dprs := CreateAeactiveDpr(keeper, ctx, 1)
	require.NotNil(t, dprs)

	keeper.DeactivateDpr(ctx, 4)

	allDprs := keeper.GetAllDpr(ctx)
	require.NotNil(t, allDprs)
}

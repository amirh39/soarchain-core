package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DeactivateDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dprs := SetupNDpr(1)
	dprsDif := SetupNDifDpr(1)
	deactiveDprs := SetupNDeactiveDpr(1)
	keeper.SetDpr(ctx, dprs[0])
	keeper.SetDpr(ctx, dprsDif[0])
	keeper.SetDpr(ctx, deactiveDprs[0])

	allDprs, found := keeper.GetAllDpr(ctx)
	require.Equal(t, found, true)
	require.NotNil(t, allDprs)

	keeper.DeactivateDpr(ctx, 4)

	allDprsAfterDeactivating := keeper.GetAllActiveDpr(ctx)
	// require.Equal(t, found, true)
	require.NotNil(t, allDprsAfterDeactivating)
}

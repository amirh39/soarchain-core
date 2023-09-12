package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_SetDprObject(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dprs := SetupNDpr(2)
	keeper.SetDpr(ctx, dprs[0])
	keeper.SetDpr(ctx, dprs[1])
	require.NotNil(t, dprs)

	dpr, found := keeper.GetDpr(ctx, dprs[0].Id)
	require.NotNil(t, dpr)
	require.Equal(t, found, true)

	allDprs := keeper.GetAllDpr(ctx)
	require.NotNil(t, allDprs)
	require.Equal(t, 2, len(allDprs))

	t.Log("allDprs", allDprs)
}

func Test_AllActiveDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dprs := SetupNDpr(2)
	deactiveDprs := SetupNDeactiveDpr(1)
	keeper.SetDpr(ctx, dprs[0])
	keeper.SetDpr(ctx, dprs[1])
	keeper.SetDpr(ctx, deactiveDprs[0])

	allDprs := keeper.GetAllDpr(ctx)
	require.NotNil(t, allDprs)

	require.Equal(t, 3, len(allDprs))

	allActiveDprs := keeper.GetAllActiveDpr(ctx)
	require.NotNil(t, allActiveDprs)

	require.Equal(t, 2, len(allActiveDprs))

}

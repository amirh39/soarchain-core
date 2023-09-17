package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_SetDprObject(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dpr := CreateDpr(keeper, ctx, 1)
	require.NotNil(t, dpr)

	got, found := keeper.GetDpr(ctx, dpr[0].Id)
	require.Equal(t, true, found)
	require.NotNil(t, got)

}

func Test_AllActiveDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dprs := CreateDeactiveDpr(keeper, ctx, 1)
	require.NotNil(t, dprs)

	allDprs := keeper.GetAllDpr(ctx)
	require.Equal(t, 1, len(allDprs))
	require.NotNil(t, allDprs)

	allActiveDprs := keeper.GetAllActiveDpr(ctx)
	require.Nil(t, allActiveDprs)
	require.Equal(t, 0, len(allActiveDprs))

}

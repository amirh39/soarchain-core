package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/testutil/nullify"
)

func Test_MasterKeyGet(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	item := CreateMasterKey(keeper, ctx)
	rst, found := keeper.GetMasterKey(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func Test_MasterKeyRemove(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	CreateMasterKey(keeper, ctx)
	keeper.RemoveMasterKey(ctx)
	_, found := keeper.GetMasterKey(ctx)
	require.False(t, found)
}

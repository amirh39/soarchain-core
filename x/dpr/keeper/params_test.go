package keeper_test

import (
	"testing"

	testkeeper "soarchain/testutil/keeper"
	"soarchain/x/dpr/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DprKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

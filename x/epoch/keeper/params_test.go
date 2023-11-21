package keeper_test

import (
	"testing"

	testkeeper "github.com/soar-robotics/soarchain-core/testutil/keeper"
	"github.com/soar-robotics/soarchain-core/x/epoch/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EpochKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

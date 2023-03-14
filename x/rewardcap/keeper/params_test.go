package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "soarchain/testutil/keeper"
	"soarchain/x/rewardcap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RewardcapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

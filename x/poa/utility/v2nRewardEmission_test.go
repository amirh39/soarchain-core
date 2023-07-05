package utility

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_V2NBXRewardEmissionPerBlock(t *testing.T) {
	ctx := sdk.Context{}
	rewardEmmission, err := V2NRewardEmissionPerBlock(ctx, "v2n-bx")
	t.Log("rewardEmmission", rewardEmmission)
	require.NotZero(t, rewardEmmission)
	require.NoError(t, err)
}

func Test_RunnerRewardEmissionPerBlock(t *testing.T) {
	ctx := sdk.Context{}
	rewardEmmission, err := V2NRewardEmissionPerBlock(ctx, "runner")
	t.Log("rewardEmmission", rewardEmmission)
	require.NotZero(t, rewardEmmission)
	require.NoError(t, err)
}

func Test_V2NBXRewardEmissionPerEpoch(t *testing.T) {
	ctx := sdk.Context{}
	rewardEmmission, err := V2NRewardEmissionPerBlock(ctx, "v2n-bx")
	t.Log("rewardEmmission", rewardEmmission)
	require.NotZero(t, rewardEmmission)
	require.NoError(t, err)
}

func Test_RunnerRewardEmissionPerEpoch(t *testing.T) {
	ctx := sdk.Context{}
	rewardEmmission, err := V2NRewardEmissionPerBlock(ctx, "runner")
	t.Log("rewardEmmission", rewardEmmission)
	require.NotZero(t, rewardEmmission)
	require.NoError(t, err)
}

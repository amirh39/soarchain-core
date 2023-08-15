package keeper_test

import (
	"fmt"
	epoch "soarchain/x/epoch/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestComputeAdaptiveHlaving(t *testing.T) {
	_, k, context, _, _, _ := SetupMsgServerForPoa(t)

	ctx := sdk.UnwrapSDKContext(context)
	// Set up some example epoch data
	initialEpochData := epoch.EpochData{
		TotalEpochs:              192,
		InitialPerChallengeValue: 100.0,
		TotalChallengesPrevDay:   50000,
	}

	updatedEpochData, _ := k.ComputeAdaptiveHalving(ctx, initialEpochData)

	expectedValueStr := "99.99999666666677456"

	assert.Equal(t, expectedValueStr, fmt.Sprintf("%.17f", updatedEpochData.InitialPerChallengeValue), "InitialPerChallengeValue doesn't match expected value")
}

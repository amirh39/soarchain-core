package utility

import (
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility/utilConstants"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_V2VRewardEmissionPerBlock(t *testing.T) {
	tests := []struct {
		currentBlockNumber     int
		expectedTokensPerBlock int
	}{
		{0, int(utilConstants.V2VReceiverInitialAnnual) / (1 * minutesPerYear)},
		{600000, int(utilConstants.V2VReceiverInitialAnnual) / (1 * minutesPerYear)},
		{2102400, int(utilConstants.V2VReceiverInitialAnnual) / (1 * minutesPerYear) / 2},
		{3679200, int(utilConstants.V2VReceiverInitialAnnual) / (1 * minutesPerYear) / 4},
	}

	for _, test := range tests {
		// Create a mock context with the current block number set to the test case value
		ctx := sdk.Context{}
		ctx = ctx.WithBlockHeight(int64(test.currentBlockNumber))

		// Call the V2VReceiveReward
		actualTokensPerBlock, err := V2VRewardEmissionPerBlock(ctx, constants.V2VRX)
		require.NoError(t, err)
		t.Log("actualTokensPerBlock", actualTokensPerBlock)
		if actualTokensPerBlock != test.expectedTokensPerBlock {
			t.Error("Test failed: expected tokens per block of, got", test.expectedTokensPerBlock, actualTokensPerBlock)
		}
	}
}

func Test_V2VRewardEmissionPerEpoch(t *testing.T) {
	ctx := sdk.Context{}
	actualTokensPerEpoch, err := V2VRewardEmissionPerEpoch(ctx, constants.V2VRX)
	t.Log("actualTokensPerEpoch", actualTokensPerEpoch)
	require.NoError(t, err)
	require.NotNil(t, actualTokensPerEpoch)
}

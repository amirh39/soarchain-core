package utility

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestV2VRewardEmissionPerBlock(t *testing.T) {
	tests := []struct {
		currentBlockNumber     int
		expectedTokensPerBlock int
	}{
		{0, 373668092100 / (1 * minutesPerYear)},
		{600000, 373668092100 / (1 * minutesPerYear)},
		{2102400, 373668092100 / (1 * minutesPerYear) / 2},
		{3679200, 373668092100 / (1 * minutesPerYear) / 4},
	}

	for _, test := range tests {
		// Create a mock context with the current block number set to the test case value
		ctx := sdk.Context{}
		ctx = ctx.WithBlockHeight(int64(test.currentBlockNumber))

		// Call the V2VReceiveReward
		actualTokensPerBlock, err := V2VRewardEmissionPerBlock(ctx, "v2v-rx")
		require.NoError(t, err)

		if actualTokensPerBlock != test.expectedTokensPerBlock {
			t.Errorf("Test failed: expected tokens per block of %f, got %f", test.expectedTokensPerBlock, actualTokensPerBlock)
		}
	}
}

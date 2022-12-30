package utility

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestV2VReceiveRewardEmissionPerBlock(t *testing.T) {
	tests := []struct {
		currentBlockNumber     int
		expectedTokensPerBlock float64
	}{
		{0, 37366809.21 / (1 * minutesPerYear)},
		{600000, 37366809.21 / (1 * minutesPerYear)},
		{2102400, 37366809.21 / (1 * minutesPerYear) / 2},
		{3679200, 37366809.21 / (1 * minutesPerYear) / 4},
	}

	for _, test := range tests {
		// Create a mock context with the current block number set to the test case value
		ctx := sdk.Context{}
		ctx = ctx.WithBlockHeight(int64(test.currentBlockNumber))

		// Call the V2VReceiveReward
		actualTokensPerBlock := V2VReceiveRewardEmissionPerBlock(ctx)

		if actualTokensPerBlock != test.expectedTokensPerBlock {
			t.Errorf("Test failed: expected tokens per block of %f, got %f", test.expectedTokensPerBlock, actualTokensPerBlock)
		}
	}
}

package utility

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CalculateRewardMultiplier(t *testing.T) {
	// Test input values
	score := CalculateRewardMultiplierScore
	expectedMultiplier := math.Pow(score, 1)

	// Call CalculateRewardMultiplier function
	multiplier := CalculateRewardMultiplier(score)

	// Check that the result is correct
	require.Equal(t, int(expectedMultiplier), int(multiplier))
}

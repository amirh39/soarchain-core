package utility

import (
	"math"
	"testing"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
)

func TestCalculateRewardMultiplier(t *testing.T) {
	// Test input values
	score := 100.0
	expectedMultiplier := math.Pow(score, 2)

	// Call CalculateRewardMultiplier function
	multiplier := CalculateRewardMultiplier(score)

	// Check that the result is correct
	require.Equal(t, int(expectedMultiplier), int(multiplier))
}

package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateRewards(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	// Set up test data
	totalAmount := 1000.0
	scores := []float64{50.0, 30.0, 20.0}

	// Execute the function to be tested
	rewards := keeper.CalculateRewards(totalAmount, scores)

	// Validate the results

	// Check if the number of rewards matches the number of scores
	require.Len(t, rewards, len(scores), "Incorrect number of rewards")

	// Check if the sum of rewards equals the total amount
	sumRewards := 0.0
	for _, reward := range rewards {
		sumRewards += reward
	}
	require.InEpsilon(t, totalAmount, sumRewards, 0.0001, "Sum of rewards does not match totalAmount")

	// Check if each reward is correctly calculated
	for i, score := range scores {
		expectedReward := totalAmount * (score / calculateTotalScore(scores))
		require.InEpsilon(t, expectedReward, rewards[i], 0.0001, "Incorrect reward calculation for score %f", score)
		t.Logf("Score: %.2f, Expected Reward: %.2f, Actual Reward: %.2f", score, expectedReward, rewards[i])
	}
}

// Helper function to calculate the total score for test validation
func calculateTotalScore(scores []float64) float64 {
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}
	return totalScore
}

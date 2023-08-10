package utility

import (
	"math/big"
	"testing"
)

func TestCalculateRewards(t *testing.T) {

	totalAmount := big.NewInt(1000)             // Example total amount
	scores := []float64{30.0, 40.0, 20.0, 10.0} // Example scores

	rewards, _ := CalculateRewards(totalAmount, scores)

	// Check if the number of rewards matches the number of scores
	if len(rewards) != len(scores) {
		t.Errorf("Expected %d rewards, but got %d", len(scores), len(rewards))
	}

	// Calculate total reward and check if it matches the total amount
	totalReward := big.NewInt(0)
	for _, reward := range rewards {
		totalReward.Add(totalReward, reward)
	}
	if totalReward.Cmp(totalAmount) != 0 {
		t.Errorf("Total reward should be equal to the total amount")
	}

	// Check if each reward is a non-negative value
	for _, reward := range rewards {
		if reward.Cmp(big.NewInt(0)) < 0 {
			t.Errorf("Reward should be non-negative")
		}
	}
}

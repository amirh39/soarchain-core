package utility

import (
	"math/big"
	"testing"
)

func TestCalculateRewards(t *testing.T) {
	// Example total amount
	totalAmount := big.NewInt(7200000)

	// Example scores and message counts for clients
	scores := []float64{50.0, 100.0, 67.0, 56.0}
	messageCounts := []int{5, 10, 8, 3}

	// Calculate expected rewards for multiple clients
	expectedRewards := []*big.Int{
		big.NewInt(1331868), // client 1
		big.NewInt(2663736), // client 2
		big.NewInt(1856703), // client 3
		big.NewInt(1347692), // client 4
	}

	// Calculate rewards using the function for multiple clients
	rewards, err := CalculateRewards(totalAmount, scores, messageCounts)

	// Check for errors
	if err != nil {
		t.Fatalf("Error calculating rewards: %v", err)
	}

	// Check if the number of rewards matches the number of clients
	if len(rewards) != len(scores) {
		t.Fatalf("Expected %d rewards, but got %d", len(scores), len(rewards))
	}

	// Check if each reward is calculated correctly for multiple clients
	for i, reward := range rewards {
		if reward.Cmp(expectedRewards[i]) != 0 {
			t.Errorf("Expected reward for client %d: %s, but got: %s", i+1, expectedRewards[i].String(), reward.String())
		}
	}

	// Example for a single client
	// Single client with the entire total amount
	singleClientScores := []float64{50.0}
	singleClientMessageCounts := []int{1}
	expectedSingleClientReward := big.NewInt(7200000)

	// Calculate rewards using the function for a single client
	singleClientRewards, err := CalculateRewards(totalAmount, singleClientScores, singleClientMessageCounts)

	// Check for errors
	if err != nil {
		t.Fatalf("Error calculating single client reward: %v", err)
	}

	// Check if the reward is calculated correctly for a single client
	if singleClientRewards[0].Cmp(expectedSingleClientReward) != 0 {
		t.Errorf("Expected reward for single client: %s, but got: %s", expectedSingleClientReward.String(), singleClientRewards[0].String())
	}
}

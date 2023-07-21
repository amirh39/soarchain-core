package keeper

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateRewards(totalAmount float64, scores []float64) []sdk.Dec {
	numScores := len(scores)

	// Calculate total score
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}

	// Calculate individual rewards concurrently in float64
	rewards := make([]float64, numScores)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(numScores)
	for i, score := range scores {
		go func(index int, s float64) {
			defer wg.Done()
			percentage := s / totalScore
			reward := totalAmount * percentage
			mu.Lock()
			rewards[index] = reward
			mu.Unlock()
		}(i, score)
	}
	wg.Wait()

	// Convert rewards to sdk.Dec
	sdkRewards := make([]sdk.Dec, numScores)
	for i, reward := range rewards {
		sdkRewards[i] = sdk.NewDecWithPrec(int64(reward*1000000), 6) // Convert to 6 decimal places
	}

	return sdkRewards
}

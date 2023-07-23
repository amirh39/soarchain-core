package keeper

import (
	"math/big"
	"sync"
)

func (k Keeper) CalculateRewards(totalAmount *big.Int, scores []float64) []*big.Int {
	numScores := len(scores)

	// Calculate total score
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}

	// Calculate individual rewards concurrently
	rewards := make([]*big.Int, numScores)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(numScores)
	for i, score := range scores {
		go func(index int, s float64) {
			defer wg.Done()
			percentage := s / totalScore
			reward := big.NewInt(int64(percentage * float64(totalAmount.Int64())))
			mu.Lock()
			rewards[index] = reward
			mu.Unlock()
		}(i, score)
	}
	wg.Wait()

	return rewards
}

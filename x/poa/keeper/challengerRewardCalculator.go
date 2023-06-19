package keeper

import "sync"

func (k Keeper) CalculateRewards(totalAmount float64, scores []float64) []float64 {
	numScores := len(scores)

	// Calculate total score
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}

	// Calculate individual rewards concurrently
	rewards := make([]float64, numScores)
	var wg sync.WaitGroup
	wg.Add(numScores)
	for i, score := range scores {
		go func(index int, s float64) {
			defer wg.Done()
			percentage := s / totalScore
			rewards[index] = totalAmount * percentage
		}(i, score)
	}
	wg.Wait()

	return rewards
}

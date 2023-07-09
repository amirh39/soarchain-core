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

	return rewards
}

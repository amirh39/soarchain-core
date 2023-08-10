package utility

import (
	"math/big"
	"sync"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CalculateRewards(totalAmount *big.Int, scores []float64) (resultRewards []*big.Int, err error) {
	if totalAmount.Cmp(big.NewInt(0)) == 0 {
		err = sdkerrors.Wrap(sdkerrors.ErrLogic, "[CalculateRewards] [totalAmount] cannot be 0")
		return
	}
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

	return rewards, nil
}

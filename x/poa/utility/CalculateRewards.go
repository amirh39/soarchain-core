package utility

import (
	"math/big"
	"sync"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CalculateRewards(totalAmount *big.Int, scores []float64, messageCounts []int) (resultRewards []*big.Int, err error) {
	if totalAmount.Cmp(big.NewInt(0)) == 0 {
		err = sdkerrors.Wrap(sdkerrors.ErrLogic, "[CalculateRewards] [totalAmount] cannot be 0")
		return
	}

	numClients := len(scores)
	if numClients != len(messageCounts) {
		err = sdkerrors.Wrap(sdkerrors.ErrLogic, "[CalculateRewards] [messageCounts] length must match [scores] length")
		return
	}

	// Calculate total score and total message count
	totalScore := 0.0
	totalMessages := 0
	for i, score := range scores {
		totalScore += score
		totalMessages += messageCounts[i]
	}

	// Calculate individual rewards concurrently
	rewards := make([]*big.Int, numClients)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(numClients)
	for i := 0; i < numClients; i++ {
		go func(index int, score float64, messageCount int) {
			defer wg.Done()
			// Calculate the reward based on both score and message count
			percentage := score / totalScore
			messagePercentage := float64(messageCount) / float64(totalMessages)
			reward := big.NewInt(int64((percentage*0.8 + messagePercentage*0.2) * float64(totalAmount.Int64())))
			mu.Lock()
			rewards[index] = reward
			mu.Unlock()
		}(i, scores[i], messageCounts[i])
	}
	wg.Wait()

	return rewards, nil
}

package keeper

func CalculateReward(totalAmount, score float64) float64 {
	// Calculate the reward amount based on the total amount of tokens and score
	reward := (totalAmount / 100.0) * score

	return reward
}

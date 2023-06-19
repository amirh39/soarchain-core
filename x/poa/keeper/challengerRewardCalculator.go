package keeper

import "math/big"

func (k Keeper) CalculateReward(totalAmount, score float64) *big.Float {
	amount := big.NewFloat(totalAmount)
	percentage := big.NewFloat(score / 100.0)

	reward := new(big.Float).Mul(amount, percentage)

	return reward
}

package utility

import (
	"math"
)

func CalculateRewardMultiplier(score float64) float64 {

	if score < 50 {
		return 0
	}

	// rewardMultiplier = (Sω)∗(1−(τ∗ψ))
	rewardMultiplier := (math.Pow(score, Omega)) * (1 - (Tau * Psi))

	return rewardMultiplier
}

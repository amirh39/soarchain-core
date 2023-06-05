package utility

import (
	"math"
	"soarchain/x/poa/utility/utilConstants"
)

func CalculateRewardMultiplier(score float64) float64 {

	if score < 50 {
		return 0
	}

	// rewardMultiplier = (Sω)∗(1−(τ∗ψ))
	rewardMultiplier := (math.Pow(score, utilConstants.Omega)) * (1 - (utilConstants.Tau * utilConstants.Psi))

	return rewardMultiplier
}

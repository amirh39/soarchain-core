package utility

import (
	"math"
)

func CalculateRewardMultiplier(score float64) float64 {

	type NetworkParams struct {
		Omega float64
		Tau   float64
		Psi   float64
	}

	// Network parameteres
	networkParams := NetworkParams{Omega: 2, Tau: 0, Psi: 0}

	if score < 50 {
		return 0
	}

	// rewardMultiplier = (Sω)∗(1−(τ∗ψ))
	rewardMultiplier := (math.Pow(score, networkParams.Omega)) * (1 - (networkParams.Tau * networkParams.Psi))

	return rewardMultiplier
}

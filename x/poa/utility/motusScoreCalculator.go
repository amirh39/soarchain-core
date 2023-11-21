package utility

import (
	"math"

	"github.com/soar-robotics/soarchain-core/x/poa/utility/utilConstants"
)

func CalculateScore(score float64, c bool) float64 {
	const minScore = utilConstants.MinScore
	const maxScore = utilConstants.MaxScore
	const maxIncreaseFactor = utilConstants.MaxIncreaseFactor
	const maxDecreaseFactor = utilConstants.MaxDecreaseFactor

	// Score should between 0 and 100
	if score < minScore || score > maxScore {
		return 0.0
	}

	increaseFactor := maxIncreaseFactor * (maxScore - score) / maxScore
	decreaseFactor := maxDecreaseFactor * score / maxScore

	if c {
		if score >= maxScore-increaseFactor {
			score += (maxScore - score) * 0.1
		} else {
			score += increaseFactor
		}
	} else {
		if score <= minScore+decreaseFactor {
			score -= (score - minScore) * 0.3
		} else {
			score -= decreaseFactor
		}
	}

	return math.Max(minScore, math.Min(maxScore, score))
}

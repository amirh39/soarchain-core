package utility

import (
	"math"
)

func CalculateScore(score float64, c bool) float64 {
	const minScore = 0
	const maxScore = 100
	const maxIncreaseFactor = 0.5
	const maxDecreaseFactor = 0.7

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

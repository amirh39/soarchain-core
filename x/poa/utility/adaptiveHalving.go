package utility

func calculateCoefficients(initialValue, targetValue float64, totalChallengesTarget int) (A, B, C float64) {
	// achieve the target value
	decayFactor := targetValue / initialValue

	// Calculate C based on the total number of challenges at the target
	C = float64(totalChallengesTarget)

	A = decayFactor*(C+1) - 1
	B = A - C

	return A, B, C
}

func calculateMintedPerChallenge(prevMintedPerChallenge float64, totalChallengesPrevDay int, A, B, C float64) float64 {
	return prevMintedPerChallenge * A / (B + float64(totalChallengesPrevDay) + C)
}

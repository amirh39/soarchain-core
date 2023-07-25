package utility

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func CalculateCoefficients(initialValue, targetValue float64, totalChallengesTarget int) (A, B, C float64) {
	// achieve the target value
	decayFactor := targetValue / initialValue

	// Calculate C based on the total number of challenges at the target
	C = float64(totalChallengesTarget)

	A = decayFactor*(C+1) - 1
	B = A - C

	return A, B, C
}

func CalculateMintedPerChallenge(prevMintedPerChallenge float64, totalChallengesPrevDay int, A, B, C float64) (float64, error) {
	// Check for division by zero
	denominator := B + float64(totalChallengesPrevDay) + C
	if denominator == 0 {
		return 0, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "division by zero in CalculateMintedPerChallenge")
	}

	mintedPerChallenge := prevMintedPerChallenge * A / denominator
	return mintedPerChallenge, nil
}

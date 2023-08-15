package utility

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CalculateCoefficients(initialValue, targetValue float64, totalChallengesTarget int) (A, B, C float64, err error) {
	// Check for division by zero
	if initialValue == 0 {
		err = sdkerrors.Wrap(sdkerrors.ErrLogic, "[CalculateCoefficients] [initialValue] cannot be 0")
		return
	}

	// Calculate the decay factor
	decayFactor := targetValue / initialValue

	// Set the value of C
	C = float64(totalChallengesTarget)

	// Calculate coefficient A using the decay factor and C
	A = decayFactor * (C)

	// Calculate coefficient B by subtracting C from A
	B = A - C

	return A, B, C, nil
}

func CalculateMintedPerChallenge(prevMintedPerChallenge float64, totalChallengesPrevDay int, A, B, C float64) (float64, error) {
	// Check for division by zero
	denominator := B + float64(totalChallengesPrevDay) + C
	if denominator == 0 {
		return 0, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[CalculateMintedPerChallenge] [denominator] division by zero in CalculateMintedPerChallenge")
	}

	// Calculate mintedPerChallenge
	mintedPerChallenge := prevMintedPerChallenge * A / denominator

	// Ensure the result is not negative
	if mintedPerChallenge < 0 {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "[CalculateMintedPerChallenge] [mintedPerChallenge] mintedPerChallenge should not be negative")
	}

	return mintedPerChallenge, nil
}

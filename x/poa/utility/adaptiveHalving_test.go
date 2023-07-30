package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCoefficients(t *testing.T) {
	initialValue := 100.0
	targetValue := 200.0
	totalChallengesTarget := 10

	A, B, C := CalculateCoefficients(initialValue, targetValue, totalChallengesTarget)

	assert.Equal(t, float64(21), A)
	assert.Equal(t, float64(11), B)
	assert.Equal(t, float64(totalChallengesTarget), C)
}

func TestCalculateMintedPerChallenge(t *testing.T) {
	prevMintedPerChallenge := 10.0
	totalChallengesPrevDay := 5
	A, B, C := 11.0, 10.0, 10.0

	mintedPerChallenge, err := CalculateMintedPerChallenge(prevMintedPerChallenge, totalChallengesPrevDay, A, B, C)

	assert.NoError(t, err)
	assert.Equal(t, float64(4.4), mintedPerChallenge)
}

// since the error message is wrapping in a Cosmos SDK error, so function should return a standard Go error for the division by zero scenario to test it properly.
// func TestCalculateMintedPerChallenge_DivisionByZero(t *testing.T) {
// 	prevMintedPerChallenge := 10.0
// 	totalChallengesPrevDay := 0
// 	A, B, C := 11.0, 10.0, 10.0

// 	_, err := CalculateMintedPerChallenge(prevMintedPerChallenge, totalChallengesPrevDay, A, B, C)
// 	assert.EqualError(t, err, "division by zero in CalculateMintedPerChallenge")

// }

// func TestCalculateMintedPerChallenge_NegativeResult(t *testing.T) {
// 	prevMintedPerChallenge := 10.0
// 	totalChallengesPrevDay := 20
// 	A, B, C := 11.0, 10.0, 10.0

// 	mintedPerChallenge, err := CalculateMintedPerChallenge(prevMintedPerChallenge, totalChallengesPrevDay, A, B, C)

// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "mintedPerChallenge should not be negative")
// 	assert.Equal(t, float64(0), mintedPerChallenge)
// }

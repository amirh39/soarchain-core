package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCoefficients(t *testing.T) {
	initialValue := 100.0
	targetValue := 200.0
	totalChallengesTarget := 10

	A, B, C, _ := CalculateCoefficients(initialValue, targetValue, totalChallengesTarget)

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

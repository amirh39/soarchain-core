package utility

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCoefficients(t *testing.T) {
	initialValue := 9.0
	targetValue := 6.0
	totalChallengesTarget := 100_000_000

	A, B, C, _ := CalculateCoefficients(initialValue, targetValue, totalChallengesTarget)

	assert.Equal(t, float64(6.6666666666666664e+07), A)
	assert.Equal(t, float64(-3.3333333333333336e+07), B)
	assert.Equal(t, float64(totalChallengesTarget), C)
}

func TestCalculateCoefficientsDeterminism(t *testing.T) {
	initialValue := 9.0
	targetValue := 6.0
	totalChallengesTarget := 100_000_000

	var differentResultsCount int

	previousA, previousB, previousC := 0.0, 0.0, 0.0

	for i := 0; i < 100_000; i++ {
		A, B, C, _ := CalculateCoefficients(initialValue, targetValue, totalChallengesTarget)

		if i > 0 && (A != previousA || B != previousB || C != previousC) {
			differentResultsCount++
		}

		previousA, previousB, previousC = A, B, C
	}

	fmt.Printf("Number of runs with different results: %d\n", differentResultsCount)

	assert.Equal(t, 0, differentResultsCount)
}
func TestCalculateMintedPerChallenge(t *testing.T) {
	prevMintedPerChallenge := 9.0
	totalChallengesPrevDay := 10_000
	A, B, C := 6.6666666666666664e+07, -3.3333333333333336e+07, 100_000_000.0

	mintedPerChallenge, err := CalculateMintedPerChallenge(prevMintedPerChallenge, totalChallengesPrevDay, A, B, C)

	assert.NoError(t, err)
	assert.Equal(t, float64(8.99865020246963), mintedPerChallenge)
}
func TestCalculateMintedPerChallengeDeterminism(t *testing.T) {
	prevMintedPerChallenge := 9.0
	totalChallengesPrevDay := 10_000
	A, B, C := 6.6666666666666664e+07, -3.3333333333333336e+07, 100_000_000.0

	var differentResultsCount int
	var mintedPerChallenge float64

	for i := 0; i < 1000; i++ {
		mintedPerChallenge, err := CalculateMintedPerChallenge(prevMintedPerChallenge, totalChallengesPrevDay, A, B, C)

		if err != nil {
			t.Errorf("Error in run %d: %v", i, err)
			continue
		}

		if mintedPerChallenge != +8.998650e+000 {
			differentResultsCount++
		}
	}

	fmt.Printf("Number of runs with different results: %d\n", differentResultsCount)

	assert.Equal(t, 0, differentResultsCount)
	assert.Equal(t, 0, mintedPerChallenge)
}

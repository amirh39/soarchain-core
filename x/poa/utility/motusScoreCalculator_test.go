package utility

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CalculateScore(t *testing.T) {
	motusScore := CalculateScore(66, true)
	require.NotZero(t, motusScore)
}

func Test_CalculateScoreLessThanMinScore(t *testing.T) {
	motusScore := CalculateScore(-1, true)
	require.Zero(t, motusScore)
}

func Test_CalculateScoreLessThanMaxScore(t *testing.T) {
	motusScore := CalculateScore(-1, true)
	require.Zero(t, motusScore)
}

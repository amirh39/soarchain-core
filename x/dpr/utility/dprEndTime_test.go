package utility

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCalculateDPREndTime(t *testing.T) {
	// Define the block time (e.g., current time)
	blockTime := time.Now()

	// Define the number of epochs for testing
	epochs := 10

	// Call the function to calculate the DPREndTime
	dprEndTimeStr, err := CalculateDPREndTime(blockTime, epochs)

	// Check for errors
	require.NoError(t, err)

	// Calculate the expected DPREndTime based on the provided number of epochs
	expectedDPREndTime := blockTime.Add(time.Duration(epochs*BlocksPerEpoch*SecondsPerBlock) * time.Second)
	expectedDPREndTimeStr := expectedDPREndTime.Format("2006-01-02 15:04 +0000 UTC")

	// Compare the calculated result with the expected result
	require.Equal(t, expectedDPREndTimeStr, dprEndTimeStr)
}

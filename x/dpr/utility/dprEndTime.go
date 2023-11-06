package utility

import (
	"time"

	"github.com/cosmos/cosmos-sdk/types/errors"
)

// CalculateDPREndTime returns the Unix time for the end of DPR based on the given epochs
func CalculateDPREndTime(blockTime time.Time, epochs int) (int64, error) {
	// Check for valid input
	if epochs < 0 {
		return 0, errors.Wrapf(errors.ErrInvalidRequest, "Negative epochs not allowed: %d", epochs)
	}

	// Calculate the duration based on the number of epochs
	duration := time.Duration(epochs*BlocksPerEpoch*SecondsPerBlock) * time.Second

	// Add the duration to the blockTime
	dprEndTime := blockTime.Add(duration)

	// Convert the end time to Unix timestamp
	dprEndTimeUnix := dprEndTime.Unix()
	return dprEndTimeUnix, nil
}

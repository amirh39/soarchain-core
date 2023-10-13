package utility

import (
	"time"

	"github.com/cosmos/cosmos-sdk/types/errors"
)

func CalculateDPREndTime(blockTime time.Time, epochs int) (string, error) {
	// Check for valid input
	if epochs < 0 {
		return "", errors.Wrapf(errors.ErrInvalidRequest, "Negative epochs not allowed: %d", epochs)
	}

	// Calculate the duration based on the number of epochs
	duration := time.Duration(epochs*BlocksPerEpoch*SecondsPerBlock) * time.Second

	// Add the duration to the blockTime
	dprEndTime := blockTime.Add(duration)
	dprEndTimeStr := dprEndTime.Format("2006-01-02 15:04 +0000 UTC")
	return dprEndTimeStr, nil
}

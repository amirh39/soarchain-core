package utility

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func IsChallengeable(ctx sdk.Context, score string, lastChallengeTime string, cooldownTolerance float64) (bool, float64, error) {

	if cooldownTolerance < 1 || cooldownTolerance > 5 {
		return false, 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "Invalid interval for cooldown tolerance parameter!")
	}

	// Convert lastChallengeTime to int
	t, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", lastChallengeTime)
	if err != nil {
		return false, 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "Can't parse lastChallengeTime!")
	}

	// Convert the time.Time value to a Unix timestamp (number of seconds since January 1, 1970 UTC).
	timestamp := t.Unix()

	// Convert the Unix timestamp to a uint.
	lastChallengeTimeInt := float64(timestamp)

	// Current time:
	currBlockTime := ctx.BlockTime()
	currentBlockTime := currBlockTime.Unix()

	// Convert the Unix timestamp to a uint.
	currentBlockTimeInt := float64(currentBlockTime)

	// Time Passed
	interval := (currentBlockTimeInt - lastChallengeTimeInt) / 60

	// Convert score to float64
	scoreFloat64, err := strconv.ParseFloat(score, 64)
	if err != nil {
		return false, 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
	}

	// Calculate challengeability
	C := (100 - scoreFloat64) + ((interval * cooldownTolerance) / 2)
	if C > 100 {
		return true, C, nil
	}
	return false, C, nil

}

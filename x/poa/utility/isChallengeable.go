package utility

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func IsChallengeable(ctx sdk.Context, score string, lastChallengeTime string, cooldownTolerance string) (bool, float64, error) {

	// Convert cooldownTolerance to uint64
	cooldownToleranceUint64, err := strconv.ParseUint(cooldownTolerance, 10, 64)
	if err != nil {
		return false, 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[IsChallengeable][ParseUint] failed. Couldn't convert to uint64. Error: [ %T ]", err)
	}

	if cooldownToleranceUint64 < 1 || cooldownToleranceUint64 > 5 {
		return false, 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[IsChallengeable] failed. Invalid interval for cooldown tolerance parameter.")
	}

	// Convert lastChallengeTime to int
	t, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", lastChallengeTime)
	if err != nil {
		return false, 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[IsChallengeable] failed. Couldn't parse lastChallengeTime. Error: [ %T ]", err)
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
		return false, 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[IsChallengeable][ParseFloat] failed. Couldn't convert score to Float64. Error: [ %T ]", err)
	}

	// Calculate challengeability
	var normalizer float64 = 60
	C := (100 - scoreFloat64) + ((interval * float64(cooldownToleranceUint64)) * normalizer)
	if C > 100 {
		return true, C, nil
	}
	return false, C, nil
}

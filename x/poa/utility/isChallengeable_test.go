package utility

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

func TestIsChallengeable(t *testing.T) {
	ctx := sdk.Context{}
	ctx = ctx.WithBlockTime(time.Date(2023, 01, 06, 11, 05, 17, 40125, time.UTC))
	// Test valid input
	isChallengeable, point, err := IsChallengeable(ctx, "80", "2022-01-06 11:05:17.40125 +0000 UTC", "1")
	require.NoError(t, err)
	require.True(t, isChallengeable)
	require.Equal(t, 105.0, point)

	// Test invalid score
	isChallengeable, point, err = IsChallengeable(ctx, "invalid", "2022-01-06 11:05:17.40125 +0000 UTC", "1")
	require.Error(t, err)
	require.Equal(t, "Cannot convert to Float64", err.Error())

	// Test invalid timeSinceLastChallenge
	isChallengeable, point, err = IsChallengeable(ctx, "80", "invalid", "1")
	require.Error(t, err)
	require.Equal(t, "Invalid interval for cooldown tolerance parameter!", err.Error())

	// Test invalid cooldownTolerance
	isChallengeable, point, err = IsChallengeable(ctx, "80", "2022-01-06 11:05:17.40125 +0000 UTC", "0")
	require.Error(t, err)
	require.Equal(t, "Invalid interval for cooldown tolerance parameter!", err.Error())

	isChallengeable, point, err = IsChallengeable(ctx, "80", "2022-01-06 11:05:17.40125 +0000 UTC", "0")
	require.Error(t, err)
	require.Equal(t, "Invalid interval for cooldown tolerance parameter!", err.Error())
}

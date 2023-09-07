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
	require.Equal(t, 3.153602e+07, point)

	// Test invalid score
	isChallengeable, point, err = IsChallengeable(ctx, "invalid", "2022-01-06 11:05:17.40125 +0000 UTC", "1")
	require.False(t, isChallengeable)
	require.NotNil(t, point)
	require.Error(t, err)

	// Test invalid timeSinceLastChallenge
	isChallengeable, point, err = IsChallengeable(ctx, "80", "invalid", "1")
	require.False(t, isChallengeable)
	require.NotNil(t, point)
	require.Error(t, err)

	// Test invalid cooldownTolerance
	isChallengeable, point, err = IsChallengeable(ctx, "80", "2022-01-06 11:05:17.40125 +0000 UTC", "0")
	require.False(t, isChallengeable)
	require.NotNil(t, point)
	require.Error(t, err)

	isChallengeable, point, err = IsChallengeable(ctx, "80", "2022-01-06 11:05:17.40125 +0000 UTC", "0")
	require.False(t, isChallengeable)
	require.NotNil(t, point)
	require.Error(t, err)
}

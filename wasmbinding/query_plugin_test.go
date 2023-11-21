package wasmbinding_test

import (
	"testing"

	"github.com/soar-robotics/soarchain-core/wasmbinding"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ChallengerByIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challenger := CreateNReputation(1)
	k.SetReputation(ctx, challenger[0])

	got, found := k.GetReputation(ctx, Challenger_PubKey)
	require.Equal(t, Score, got.Score)
	require.Equal(t, true, found)
	require.NotEmpty(t, 1, got)
}

func Test_ChallengerByNotValidIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challenger := CreateNReputation(1)
	k.SetReputation(ctx, challenger[0])

	client, err := wasmbinding.GetChallenger(ctx, NotValidndex, k)
	require.Error(t, err)
	require.Empty(t, client)
}

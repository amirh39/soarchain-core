package wasmbinding_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"soarchain/wasmbinding"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestChallengerByIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challengers := CreateNReputation(1)
	k.SetReputation(ctx, challengers[0])

	challenger, err := wasmbinding.GetChallenger(ctx, Challenger_PubKey, k)
	require.NotEmpty(t, 1, challenger)
	require.Equal(t, "189", challenger.Score)
	require.Equal(t, nil, err)
}

func TestChallengerByNotValidIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challengers := CreateNReputation(1)
	k.SetReputation(ctx, challengers[0])

	challenger, err := wasmbinding.GetChallenger(ctx, NotValidndex, k)
	require.Error(t, err)
	require.Empty(t, challenger)
}

package wasmbinding_test

import (
	"soarchain/wasmbinding"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ChallengerByIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challenger := CreateNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	got, found := k.GetChallenger(ctx, Challenger_Address)

	require.Equal(t, Score, got.Score)
	require.Equal(t, true, found)
	require.NotEmpty(t, 1, got)
}

func Test_ChallengerByNotValidIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challenger := CreateNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	client, err := wasmbinding.GetChallenger(ctx, NotValidndex, k)
	require.Error(t, err)
	require.Empty(t, client)
}

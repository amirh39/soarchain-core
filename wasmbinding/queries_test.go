package wasmbinding_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"soarchain/wasmbinding"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestClientByIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challengers := CreateNChallenger(1)
	k.SetChallenger(ctx, challengers[0])

	challenger, err := wasmbinding.GetChallenger(ctx, Challenger_Address, k)
	require.NotEmpty(t, 1, challenger)
	require.Equal(t, "189", challenger.Score)
	require.Equal(t, nil, err)
}

func TestClientByNotValidIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	challengers := CreateNChallenger(1)
	k.SetChallenger(ctx, challengers[0])

	client, err := wasmbinding.GetChallenger(ctx, NotValidndex, k)
	require.Error(t, err)
	require.Empty(t, client)
}

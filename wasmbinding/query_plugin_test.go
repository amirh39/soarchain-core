package wasmbinding_test

import (
	"soarchain/wasmbinding"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ClientByIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	clients := SetupClientEntity(1)
	require.NotEmpty(t, 1, clients)
	k.SetClient(ctx, clients[0])

	client, found := k.GetClient(ctx, Index)

	require.Equal(t, Score, client.Score)
	require.Equal(t, true, found)
	require.NotEmpty(t, 1, client)
}

func Test_ClientByNotValidIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	clients := SetupClientEntity(1)
	k.SetClient(ctx, clients[0])

	client, err := wasmbinding.GetClientByIndex(ctx, NotValidndex, k)
	require.Error(t, err)
	require.Empty(t, client)
}

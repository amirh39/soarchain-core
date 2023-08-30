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

	clients := SetupClientEntity(1)
	require.NotEmpty(t, 1, clients)
	k.SetClient(ctx, clients[0])

	client, err := wasmbinding.GetClientByIndex(ctx, Index, k)
	require.NotEmpty(t, 1, client)
	require.Equal(t, "25", client.Score)
	require.Equal(t, nil, err)
}

func TestClientByNotValidIndex(t *testing.T) {
	k, context := SetupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	clients := SetupClientEntity(1)
	k.SetClient(ctx, clients[0])

	client, err := wasmbinding.GetClientByIndex(ctx, NotValidndex, k)
	require.Error(t, err)
	require.Empty(t, client)
}

package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenClient(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	clients := SetupClientEntity(1)
	k.SetClient(ctx, clients[0])

	motusWallet := SetupMotusWalletEntityByClient(clients[0])
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     ADDRESS,
		Certificate: CERTIFICATE,
		Signature:   Signature,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

func Test_GenClient_NotValidCertificate(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	clients := SetupClientEntity(1)
	k.SetClient(ctx, clients[0])

	motusWallet := SetupMotusWalletEntityByClient(clients[0])
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     ADDRESS,
		Certificate: NOTVALIDCERTIFICATE,
		Signature:   Signature,
	})

	require.Error(t, err)
	require.Nil(t, res)
}

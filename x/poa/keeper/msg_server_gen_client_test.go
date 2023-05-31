package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenClient(t *testing.T) {
	msgServer, _, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	keeper, ctx := keepertest.PoaKeeper(t)

	clients := CreateGenClient(keeper, ctx, 1)
	require.NotNil(t, clients)

	motusWallet := CreateMotusWalletByClientEntity(keeper, ctx, clients[0])
	require.NotNil(t, motusWallet)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     ADDRESS,
		Certificate: CERTIFICATE,
		Signature:   Signature,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

func Test_GenClient_NotValidCertificate(t *testing.T) {
	msgServer, _, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	keeper, ctx := keepertest.PoaKeeper(t)

	clients := CreateGenClient(keeper, ctx, 1)

	CreateMotusWalletByClientEntity(keeper, ctx, clients[0])

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     ADDRESS,
		Certificate: NOTVALIDCERTIFICATE,
		Signature:   Signature,
	})

	require.Error(t, err)
	require.Nil(t, res)
}

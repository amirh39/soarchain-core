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

	item := types.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	registeredFactoryKey, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
		Creator:     CREATOR,
		FactoryCert: FACTORY_CERT,
	})
	require.Nil(t, err)
	require.NoError(t, err)
	require.NotNil(t, registeredFactoryKey)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     ClientForTestUniqe,
		Certificate: CERTIFICATE1,
		Signature:   SIGNATURE1,
	})

	clients := k.GetAllClient(ctx)
	t.Log("clients------------------->", clients)

	require.NotNil(t, res)
	require.NoError(t, err)
}

func Test_GenClient1(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	registeredFactoryKey, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
		Creator:     CREATOR,
		FactoryCert: FACTORY_CERT,
	})
	require.Nil(t, err)
	require.NoError(t, err)
	require.NotNil(t, registeredFactoryKey)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     ClientForTestUniqe,
		Certificate: CERTIFICATE2,
		Signature:   SIGNATURE2,
	})

	clients := k.GetAllClient(ctx)
	t.Log("clients------------------->", clients)

	require.NotNil(t, res)
	require.NoError(t, err)
	require.NoError(t, err)
}

func Test_UniqueClient(t *testing.T) {

	_, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	Test_GenClient(t)
	Test_GenClient1(t)

	clients := k.GetAllClient(ctx)
	t.Log("clients------------->", clients)

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

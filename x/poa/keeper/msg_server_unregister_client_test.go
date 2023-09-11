package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterClient(t *testing.T) {
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
	require.NoError(t, err)
	require.NotNil(t, registeredFactoryKey)

	clients := SetupClientToUnregistration(1)
	k.SetClient(ctx, clients[0])

	motusWallet := SetupMotusWalletEntityByClient(clients[0])
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.UnregisterClient(context, &types.MsgUnregisterClient{
		Creator: CREATOR,
		Pubkey:  ClientPubKey,
	})
	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid client, response should raise proper error message*/
func Test_UnregisterChallenger_NotValidClient(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	client := SetupClientWithInvalidScore(1)

	motusWallet := SetupMotusWalletEntityByClient(client[0])
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.UnregisterClient(context, &types.MsgUnregisterClient{
		Creator: ClientAddress,
		Pubkey:  ClientPubKey,
	})
	require.Error(t, err)
	require.Nil(t, res)
}

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

	client := SetupClientEntity(1)
	k.SetClient(ctx, client[0])

	motusWallet := SetupMotusWalletEntityByClient(client[0])
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.UnregisterClient(context, &types.MsgUnregisterClient{
		Creator: ClientAddress,
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

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"soarchain/x/poa/types"
)

func Test_ClaimMotusRewards(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)
	// create test motus wallet and client
	motusWallet := types.MotusWallet{
		Address: CommunityWallet,
		Client:  &types.Client{Address: CommunityWallet, NetEarnings: MotusWalletAmount},
	}
	k.SetMotusWallet(ctx, motusWallet)

	// test with successful claim
	resp, err := msgServer.ClaimMotusRewards(context, &types.MsgClaimMotusRewards{
		Creator: CommunityWallet,
		Amount:  MotusWalletAmount,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

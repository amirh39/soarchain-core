package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Challenger{
		PubKey:       "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251",
		Address:      "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
		Score:        "189",
		StakedAmount: "2000000000utmotus",
		NetEarnings:  "0utmotus",
		IpAddr:       "",
		Type:         "v2n",
	}
	k.SetChallenger(ctx, item)

	res, err := msgServer.UnregisterChallenger(context, &types.MsgUnregisterChallenger{
		Creator:           "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
		ChallengerAddress: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
	})

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid guard, response should raise proper error message*/
func Test_UnregisterChallenger_NotValidGuard(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Challenger{
		PubKey:       "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251",
		Address:      "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
		Score:        "189",
		StakedAmount: "2000000000utmotus",
		NetEarnings:  "0utmotus",
		IpAddr:       "",
		Type:         "v2n",
	}
	k.SetChallenger(ctx, item)

	res, err := msgServer.UnregisterChallenger(context, &types.MsgUnregisterChallenger{
		Creator:           "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
		ChallengerAddress: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

/** Using not valid challenger, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_NotValidChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Challenger{
		PubKey:       "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251",
		Address:      "",
		Score:        "189",
		StakedAmount: "2000000000utmotus",
		NetEarnings:  "0utmotus",
		IpAddr:       "",
		Type:         "v2n",
	}
	k.SetChallenger(ctx, item)

	res, err := msgServer.UnregisterChallenger(context, &types.MsgUnregisterChallenger{
		Creator:           "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
		ChallengerAddress: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	resp, err := msgServer.GenChallenger(context, &types.MsgGenChallenger{
		Creator:          CREATOR,
		ChallengerPubKey: "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251",
		ChallengerAddr:   "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y",
		ChallengerStake:  "2000000000utmotus",
		ChallengerIp:     "104.248.142.45",
		Challengertype:   "v2n",
		Certificate:      CERTIFICATE,
		Signature:        "",
	})

	t.Log("response", resp)

	require.NoError(t, err)
	require.NotNil(t, resp)
}

/** Using not valid master certificate, response should raise proper error message*/
// func Test_GenRunner_NotValidCertificate(t *testing.T) {
// 	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
// 	defer ctrl.Finish()

// 	bank.ExpectAny(context)

// 	ctx := sdk.UnwrapSDKContext(context)

// 	item := types.MasterKey{MasterCertificate: MASTER_ACCOUNT,
// 		MasterAccount: MASTER_ACCOUNT,
// 	}
// 	k.SetMasterKey(ctx, item)

// 	res, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
// 		Creator:     CREATOR,
// 		FactoryCert: FACTORY_CERT,
// 	})

// 	t.Log("error message", err)

// 	require.Error(t, err)
// 	require.Nil(t, res)
// }

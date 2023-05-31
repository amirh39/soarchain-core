package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.GenChallenger(context, &types.MsgGenChallenger{
		Creator:         CREATOR,
		ChallengerStake: Challenger_StakedAmount,
		ChallengerIp:    Challenger_IPAddress,
		Challengertype:  Challenger_Type,
		Certificate:     CERTIFICATE,
		Signature:       Signature,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid certificate, response should raise proper error message*/
func Test_GenChallenger_NotValidCertificate(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_ACCOUNT,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.GenChallenger(context, &types.MsgGenChallenger{
		Creator:         CREATOR,
		ChallengerStake: Challenger_StakedAmount,
		ChallengerIp:    Challenger_IPAddress,
		Challengertype:  Challenger_Type,
		Certificate:     INValid_CertString,
		Signature:       Signature,
	})

	require.Error(t, err)
	require.Nil(t, res)
}

package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenRunner(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.GenRunner(context, &types.MsgGenRunner{
		Creator:     CREATOR,
		RunnerStake: STAKE,
		RunnerIp:    IP,
		Certificate: CERTIFICATE,
		Signature:   SIGNATURE,
	})

	t.Log("response", res)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid certificate, response should raise proper error message*/
func Test_GenRunner_NotValidCertificate(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_ACCOUNT,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.GenRunner(context, &types.MsgGenRunner{
		Creator:     CREATOR,
		RunnerStake: STAKE,
		RunnerIp:    IP,
		Certificate: NOTVALIDCERTIFICATE,
		Signature:   SIGNATURE,
	})

	t.Log("error message", err)

	require.Error(t, err)
	require.Nil(t, res)
}

package keeper_test

import (
	"testing"

	"github.com/soar-robotics/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_RegisterFactoryKey(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	resp, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
		Creator:     CREATOR,
		FactoryCert: FACTORY_CERT,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
}

/** Using not valid master certificate, response should raise proper error message*/
func Test_RegisterNotValidFactoryKey(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_ACCOUNT,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
		Creator:     CREATOR,
		FactoryCert: FACTORY_CERT,
	})

	require.Error(t, err)
	require.Nil(t, res)
}

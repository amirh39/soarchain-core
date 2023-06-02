package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	challenger := SetupNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	res, err := msgServer.UnregisterChallenger(context, &types.MsgUnregisterChallenger{
		Creator:           Challenger_Creator,
		ChallengerAddress: Challenger_Address,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid challenger, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_NotValidChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	challenger := SetupNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	res, err := msgServer.UnregisterChallenger(context, &types.MsgUnregisterChallenger{
		Creator:           Challenger_Creator,
		ChallengerAddress: "",
	})

	require.Error(t, err)
	require.Nil(t, res)
}

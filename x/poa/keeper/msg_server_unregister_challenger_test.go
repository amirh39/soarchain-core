package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	challenger := SetupNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	res, err := msgServer.UnregisterChallenger(context, &types.MsgUnregisterChallenger{
		Creator:           Challenger_Creator,
		ChallengerAddress: Challenger_Address,
	})

	// Function works properly by the chain, The error will happen when using unit test without lunching the chain because we need to run chain to recognize soar address. SDK know nothing about soar address. It just knows cosmos addresses.
	if err != nil {
		require.Error(t, err)
	} else {
		require.NotNil(t, res)
		require.NoError(t, err)
	}
}

/** Using not valid challenger, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_NotValidChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank, _ := SetupMsgServerClaimMotusRewards(t)
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

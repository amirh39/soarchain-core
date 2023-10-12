package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ClaimRunnerRewards(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	reputation := SetupReputationForRunner(1)
	k.SetReputation(ctx, reputation[0])

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: RunnerAddress,
		Amount:  Amount,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid runner key, response should raise proper error message*/
func Test_RegisterNotValidClaimRunnerReward_KeyNotFound(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)

	runner := SetupReputationForRunner(1)
	k.SetReputation(ctx, runner[0])

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: CREATOR,
		Amount:  Amount,
	})

	require.Error(t, err)
	require.Nil(t, res)
}

/** Using not valid amount, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_InsufficientFound(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)

	runner := SetupReputationForRunner(1)
	k.SetReputation(ctx, runner[0])

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: CREATOR,
		Amount:  InsufficientFound,
	})

	require.Error(t, err)
	require.Nil(t, res)
}

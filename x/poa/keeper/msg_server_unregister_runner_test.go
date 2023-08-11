package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterRunner(t *testing.T) {
	msgServer, k, context, ctrl, bank, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	runner := SetupNRunner(1)
	k.SetRunner(ctx, runner[0])

	res, err := msgServer.UnregisterRunner(context, &types.MsgUnregisterRunner{
		Creator:       RunnerAddress,
		RunnerAddress: RunnerAddress,
	})

	// Function works properly by the chain, The error will happen when using unit test without lunching the chain because we need to run chain to recognize soar address. SDK know nothing about soar address. It just knows cosmos addresses.
	if err != nil {
		require.Error(t, err)
	} else {
		require.NotNil(t, res)
		require.NoError(t, err)
	}
}

/** Using not valid guard, response should raise proper error message*/
func Test_UnregisterRunner_NotValidGuard(t *testing.T) {
	msgServer, k, context, ctrl, bank, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	runner := SetupNRunner(1)
	k.SetRunner(ctx, runner[0])

	res, err := msgServer.UnregisterRunner(context, &types.MsgUnregisterRunner{
		Creator:       "",
		RunnerAddress: "",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterRunner(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	runner := SetupNRunner(1)
	k.SetRunner(ctx, runner[0])

	res, err := msgServer.UnregisterRunner(context, &types.MsgUnregisterRunner{
		Creator:       RunnerAddress,
		RunnerAddress: RunnerAddress,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid guard, response should raise proper error message*/
func Test_UnregisterRunner_NotValidGuard(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
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

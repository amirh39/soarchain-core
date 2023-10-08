package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_SelectRandomRunner(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	runnerReputation := SetupNRunnerReputation(1)
	k.SetReputation(ctx, runnerReputation[0])

	res, err := msgServer.SelectRandomRunner(context, &types.MsgSelectRandomRunner{
		Creator: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid runner, response should raise proper error message*/
func Test_SelectRandomRunner_NotValidCreator(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	runner := SetupNRunnerReputation(1)
	k.SetReputation(ctx, runner[0])

	res, err := msgServer.SelectRandomRunner(context, &types.MsgSelectRandomRunner{
		Creator: "",
	})

	require.Error(t, err)
	require.Nil(t, res)
}

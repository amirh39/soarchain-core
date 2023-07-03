package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_RunnerChallengey(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	client := SetupClientEntity(1)
	k.SetClient(ctx, client[0])

	runner := SetupNRunner(1)
	k.SetRunner(ctx, runner[0])

	challenger := SetupNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	clientPubkeys := []string{client[0].PubKey}

	resp, err := msgServer.RunnerChallenge(context, &types.MsgRunnerChallenge{
		Creator:         Challenger_Address,
		RunnerpubKey:    runner[0].PubKey,
		ClientPubkeys:   clientPubkeys,
		ChallengeResult: "punish",
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
}

package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestRunnerChallenge(t *testing.T) {

	// Create a new instance of the msgServer
	msgServer, k, context, ctrl, _, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()
	ctx := sdk.UnwrapSDKContext(context)
	client := SetupClientEntity(1)
	k.SetClient(ctx, client[0])
	clientPubkeys := []string{client[0].Index}

	// Create a new message for the RunnerChallenge
	msg := &types.MsgRunnerChallenge{
		Creator:         Challenger_Address,
		RunnerpubKey:    RunnerPubKey,
		ClientPubkeys:   clientPubkeys,
		ChallengeResult: "reward",
	}

	// Call the RunnerChallenge function
	res, err := msgServer.RunnerChallenge(context, msg)

	// Check if there was an error
	require.NoError(t, err)

	// Check if the response is not nil
	require.NotNil(t, res)
}

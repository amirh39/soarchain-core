package keeper_test

import (
	k "soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ClaimChallengerRewards(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	ctx := sdk.UnwrapSDKContext(context)
	bank.ExpectAny(context)

	reputation := types.Reputation{
		PubKey:      Challenger_PubKey,
		Address:     Challenger_Address,
		Score:       Challenger_Score,
		NetEarnings: "100udmotus",
		Type:        Challenger_Type,
	}
	k.SetReputation(ctx, reputation)

	// Create a test message
	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "50udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	updatedChallenger, found := k.GetReputation(ctx, Challenger_PubKey)
	require.True(t, found)
	require.Equal(t, "50udmotus", updatedChallenger.NetEarnings)
}

/** Using an unknown challenger key should raise an error */
func Test_ClaimChallengerRewards_KeyNotFound(t *testing.T) {
	msgServer, _, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "100udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
}

/** Using an insufficient funds amount should raise an error */
func Test_ClaimChallengerRewards_InsufficientFunds(t *testing.T) {
	msgServer, k, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()
	ctx := sdk.UnwrapSDKContext(context)

	reputation := types.Reputation{
		PubKey:      Challenger_PubKey,
		Address:     Challenger_Address,
		Score:       Challenger_Score,
		NetEarnings: "100udmotus",
		Type:        Challenger_Type,
	}
	k.SetReputation(ctx, reputation)

	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "1000udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
}

func Test_ClaimChallengerRewards_InvalidChallengerType(t *testing.T) {
	msgServer, k, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()
	ctx := sdk.UnwrapSDKContext(context)

	reputation := types.Reputation{
		PubKey:      Challenger_PubKey,
		Address:     Challenger_Address,
		Score:       Challenger_Score,
		NetEarnings: "100udmotus",
		Type:        "InvalidType",
	}
	k.SetReputation(ctx, reputation)

	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "50udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
}


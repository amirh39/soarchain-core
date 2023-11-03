package keeper_test

import (
	k "soarchain/x/poa/keeper"
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
func Test_ClaimRunnerReward_KeyNotFound(t *testing.T) {
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
func Test_ClaimRunnerReward_InsufficientFound(t *testing.T) {
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

func Test_ClaimRunnerRewards_InvalidNetEarningsFormat(t *testing.T) {
	msgServer, k, context, _, _ := SetupMsgServerClaimMotusRewards(t)
	ctx := sdk.UnwrapSDKContext(context)

	// Set up a runner reputation with invalid net earnings format
	reputation := types.Reputation{
		PubKey:      RunnerPubKey,
		Address:     RunnerAddress,
		Score:       RunnerScore,
		NetEarnings: "invalidformat",
		Type:        RunnerType,
	}
	k.SetReputation(ctx, reputation)

	// Try to withdraw with a valid amount
	msg := types.NewMsgClaimRunnerRewards(RunnerAddress, "10udmotus")
	res, err := msgServer.ClaimRunnerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
	require.Contains(t, err.Error(), "couldn't be parsed")
}

func (helper *KeeperTestHelper) Test_ClaimRunnerRewards_InvalidNetEarningsFormat_AppLevel() {
	helper.Run("Test_ClaimRunnerRewards_InvalidNetEarningsFormat_AppLevel", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		bankKeeper := helper.App.BankKeeper

		// Mint coins to ensure the module account has enough balance
		parsedCoin, _ := sdk.ParseCoinNormalized("10000udmotus")
		bankKeeper.MintCoins(helper.Ctx, "poa", sdk.Coins{parsedCoin})

		// Set up the runner's reputation with an invalid format for net earnings
		CreateNRunnerReputation(&keeper, helper.Ctx, 1)

		runner, _ := keeper.GetReputation(helper.Ctx, RunnerPubKey)
		runner.NetEarnings = "invalidformat"
		keeper.SetReputation(helper.Ctx, runner)

		// Attempt to claim rewards with a valid amount
		msgServer := k.NewMsgServerImpl(keeper)
		msg := types.NewMsgClaimRunnerRewards(RunnerAddress, "10udmotus")
		res, err := msgServer.ClaimRunnerRewards(sdk.WrapSDKContext(helper.Ctx), msg)
		helper.Error(err)
		helper.Nil(res)
	})
}

func Test_ClaimRunnerRewards_FullWithdrawal(t *testing.T) {
	msgServer, k, context, _, bank := SetupMsgServerClaimMotusRewards(t)
	ctx := sdk.UnwrapSDKContext(context)

	bank.ExpectAny(context)

	// Set up a runner reputation
	reputation := types.Reputation{
		PubKey:      RunnerPubKey,
		Address:     RunnerAddress,
		Score:       RunnerScore,
		NetEarnings: "100udmotus",
		Type:        RunnerType,
	}
	k.SetReputation(ctx, reputation)

	// Withdraw the full amount of net earnings
	msg := types.NewMsgClaimRunnerRewards(RunnerAddress, "100udmotus")
	res, err := msgServer.ClaimRunnerRewards(context, msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	// Check if net earnings is updated to zero
	updatedRunner, found := k.GetReputation(ctx, RunnerPubKey)
	require.True(t, found)
	require.Equal(t, "0udmotus", updatedRunner.NetEarnings)
}

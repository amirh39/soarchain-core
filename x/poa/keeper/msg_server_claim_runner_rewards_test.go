package keeper_test

import (
	"strings"
	"testing"

	k "github.com/amirh39/soarchain-core/x/poa/keeper"
	"github.com/amirh39/soarchain-core/x/poa/types"

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

func Test_ClaimRunnerRewards_NotRegisteredAsRunner(t *testing.T) {
	msgServer, k, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	ctx := sdk.UnwrapSDKContext(context)

	// Create a reputation for an address with a non-runner type
	nonRunnerReputation := types.Reputation{
		Address: Challenger_Address,
		Type:    Challenger_Type,
	}
	k.SetReputation(ctx, nonRunnerReputation)

	// Try to claim runner rewards with an address not registered as a runner
	msg := &types.MsgClaimRunnerRewards{
		Creator: Challenger_Address,
		Amount:  "10udmotus",
	}

	res, err := msgServer.ClaimRunnerRewards(context, msg)

	require.Error(t, err, "Expected error when non-runner address attempts to claim runner rewards")
	require.True(t, strings.Contains(err.Error(), "not registered"), "Error message should indicate the address is not registered as a runner")
	require.Nil(t, res, "Result should be nil when claiming is unauthorized")
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

func (helper *KeeperTestHelper) Test_ClaimRunnerRewards_FullWithdrawal_AppLevel() {
	helper.Run("Test_ClaimRunnerRewards_FullWithdrawal_AppLevel", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		bankKeeper := helper.App.BankKeeper

		// Mint coins to ensure the module account has enough balance to cover the withdrawal
		parsedCoin, _ := sdk.ParseCoinNormalized("10000udmotus")
		bankKeeper.MintCoins(helper.Ctx, "poa", sdk.Coins{parsedCoin})

		// Set up the runner's reputation with net earnings that will be fully withdrawn
		CreateNRunnerReputation(&keeper, helper.Ctx, 1)

		runner, _ := keeper.GetReputation(helper.Ctx, RunnerPubKey)
		runner.NetEarnings = "100udmotus"
		keeper.SetReputation(helper.Ctx, runner)

		// Use the msgServer to handle the claim rewards message
		msgServer := k.NewMsgServerImpl(keeper)
		msg := types.NewMsgClaimRunnerRewards(RunnerAddress, "100udmotus")
		res, err := msgServer.ClaimRunnerRewards(sdk.WrapSDKContext(helper.Ctx), msg)
		helper.NoError(err)
		helper.NotNil(res)

		// Check if the runner's net earnings have been updated to zero
		updatedRunner, found := keeper.GetReputation(helper.Ctx, RunnerPubKey)
		helper.True(found)
		helper.Equal("0udmotus", updatedRunner.NetEarnings)
	})
}

func (helper *KeeperTestHelper) Test_ClaimRunnerRewards_MultipleWithdrawals_AppLevel() {
	helper.Run("Test_ClaimRunnerRewards_MultipleWithdrawals_AppLevel", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		bankKeeper := helper.App.BankKeeper

		initialBalance, _ := sdk.ParseCoinNormalized("10000udmotus")
		// Mint coins to ensure the module account has enough balance
		bankKeeper.MintCoins(helper.Ctx, types.ModuleName, sdk.Coins{initialBalance})

		// Set up the runner's reputation with initial net earnings
		runnerReputation := CreateNRunnerReputation(&keeper, helper.Ctx, 1)
		runnerReputation[0].NetEarnings = initialBalance.String()
		keeper.SetReputation(helper.Ctx, runnerReputation[0])

		withdrawalAmounts := []string{"200udmotus", "50udmotus", "1000udmotus"}

		// Perform the withdrawals and check the balances after each withdrawal
		for _, amount := range withdrawalAmounts {
			msg := types.NewMsgClaimRunnerRewards(RunnerAddress, amount)
			msgServer := k.NewMsgServerImpl(keeper)
			_, err := msgServer.ClaimRunnerRewards(sdk.WrapSDKContext(helper.Ctx), msg)
			helper.NoError(err)

			withdrawalCoin, _ := sdk.ParseCoinNormalized(amount)
			initialBalance = initialBalance.Sub(withdrawalCoin)

			updatedRunner, found := keeper.GetReputation(helper.Ctx, runnerReputation[0].PubKey)
			helper.True(found)

			updatedBalance, _ := sdk.ParseCoinNormalized(updatedRunner.NetEarnings)

			helper.Equal(initialBalance, updatedBalance)
		}

		finalRunner, found := keeper.GetReputation(helper.Ctx, runnerReputation[0].PubKey)
		helper.True(found)

		finalBalance, _ := sdk.ParseCoinNormalized(finalRunner.NetEarnings)

		helper.Equal(initialBalance, finalBalance)
	})
}

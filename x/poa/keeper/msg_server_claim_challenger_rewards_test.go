package keeper_test

import (
	"testing"

	k "github.com/amirh39/soarchain-core/x/poa/keeper"
	"github.com/amirh39/soarchain-core/x/poa/types"

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

func Test_ClaimChallengerRewards_InvalidAmountFormat(t *testing.T) {
	msgServer, _, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "abcudmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
}

func (helper *KeeperTestHelper) Test_ClaimChallengerRewards_ZeroWithdrawAmount() {
	helper.Run("Test_ClaimChallengerRewards_ZeroWithdrawAmount", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper

		CreateNChallengerReputationWithNormalScore(&keeper, helper.Ctx, 1)

		msgServer := k.NewMsgServerImpl(keeper)

		msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "0udmotus")
		res, err := msgServer.ClaimChallengerRewards(sdk.WrapSDKContext(helper.Ctx), msg)
		helper.NoError(err)
		helper.Empty(res)

		updatedChallenger, found := keeper.GetReputation(helper.Ctx, Challenger_PubKey)
		helper.True(found)
		helper.Equal("0udmotus", updatedChallenger.NetEarnings)
	})
}

func Test_ClaimChallengerRewards_EmptyNetEarnings(t *testing.T) {
	msgServer, k, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()
	ctx := sdk.UnwrapSDKContext(context)

	reputation := types.Reputation{
		PubKey:      Challenger_PubKey,
		Address:     Challenger_Address,
		Score:       Challenger_Score,
		NetEarnings: "",
		Type:        Challenger_Type,
	}
	k.SetReputation(ctx, reputation)

	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "50udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
}

func Test_ClaimChallengerRewards_InvalidChallengerAddress(t *testing.T) {
	msgServer, _, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	invalidAddress := "InvalidAddress"
	msg := types.NewMsgClaimChallengerRewards(invalidAddress, "50udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.Error(t, err)
	require.Nil(t, res)
}

func Test_ClaimChallengerRewards_FullWithdrawal(t *testing.T) {
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

	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "100udmotus")
	res, err := msgServer.ClaimChallengerRewards(context, msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	updatedChallenger, found := k.GetReputation(ctx, Challenger_PubKey)
	require.True(t, found)
	require.Equal(t, "0udmotus", updatedChallenger.NetEarnings)
}

func (helper *KeeperTestHelper) Test_ClaimChallengerRewards_FullWithdrawal_AppLevel() {
	helper.Run("Test_ClaimChallengerRewards_FullWithdrawal_AppLevel", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		bankKeeper := helper.App.BankKeeper
		parsedCoin, err := sdk.ParseCoinNormalized("10000udmotus")
		bankKeeper.MintCoins(helper.Ctx, "poa", sdk.Coins{parsedCoin})

		CreateNChallengerReputationWithNormalScore(&keeper, helper.Ctx, 1)

		challenger, _ := keeper.GetReputation(helper.Ctx, Challenger_PubKey)
		challenger.NetEarnings = "100udmotus"
		keeper.SetReputation(helper.Ctx, challenger)

		msgServer := k.NewMsgServerImpl(keeper)

		msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "100udmotus")
		res, err := msgServer.ClaimChallengerRewards(sdk.WrapSDKContext(helper.Ctx), msg)
		helper.NoError(err)
		helper.Empty(res)

		updatedChallenger, found := keeper.GetReputation(helper.Ctx, Challenger_PubKey)
		helper.True(found)
		helper.Equal("0udmotus", updatedChallenger.NetEarnings)

	})
}

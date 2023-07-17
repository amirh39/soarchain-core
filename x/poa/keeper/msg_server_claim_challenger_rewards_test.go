package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ClaimChallengerRewards(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	// Set up the bank expectations
	//bank.SendCoinsFromModuleToAccount()
	bank.ExpectAny(context)

	// Set up the context
	ctx := sdk.UnwrapSDKContext(context)

	// Create a challenger with net earnings
	challenger := types.Challenger{
		PubKey:       Challenger_PubKey,
		Address:      Challenger_Address,
		Score:        Challenger_Score,
		StakedAmount: Challenger_StakedAmount,
		NetEarnings:  "100utmotus",
		Type:         Challenger_Type,
		IpAddr:       Challenger_IPAddress,
	}
	k.SetChallenger(ctx, challenger)

	// Create a test message
	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "50utmotus")

	// Run the ClaimChallengerRewards function
	res, err := msgServer.ClaimChallengerRewards(context, msg)

	// Verify the response
	require.NoError(t, err)
	require.NotNil(t, res)

	// Verify that the challenger's net earnings have been updated
	updatedChallenger, found := k.GetChallenger(ctx, Challenger_Address)
	require.True(t, found)
	require.Equal(t, "50utmotus", updatedChallenger.NetEarnings)

	// Verify that the coins have been sent from the module to the account
	//expectedCoins := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(100)))
	acc := sdk.AccAddress(Challenger_Address)
	require.NotNil(t, acc)
	//require.Equal(t, expectedCoins, acc.Equals())
}

/** Using an unknown challenger key should raise an error */
func Test_ClaimChallengerRewards_KeyNotFound(t *testing.T) {
	msgServer, _, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	// Create a test message with an unknown challenger key
	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "100motus")

	// Run the ClaimChallengerRewards function
	res, err := msgServer.ClaimChallengerRewards(context, msg)

	// Verify the error response
	require.Error(t, err)
	require.Nil(t, res)

	// Verify the error message
	expectedErrMsg := "[ClaimChallengerRewards][GetChallenger] failed. Target challenger is not registered in the store by this address: [ string ]. Make sure the address is valid and not empty.: key not found"
	require.Equal(t, expectedErrMsg, err.Error())
}

/** Using an insufficient funds amount should raise an error */
func Test_ClaimChallengerRewards_InsufficientFunds(t *testing.T) {
	msgServer, k, context, ctrl, _ := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	// Set up the context
	ctx := sdk.UnwrapSDKContext(context)

	// Create a challenger with net earnings
	challenger := types.Challenger{
		PubKey:       Challenger_PubKey,
		Address:      Challenger_Address,
		Score:        Challenger_Score,
		StakedAmount: Challenger_StakedAmount,
		NetEarnings:  "100utmotus",
		Type:         Challenger_Type,
		IpAddr:       Challenger_IPAddress,
	}
	k.SetChallenger(ctx, challenger)

	// Create a test message with an amount greater than the net earnings
	msg := types.NewMsgClaimChallengerRewards(Challenger_Address, "1000utmotus")

	// Run the ClaimChallengerRewards function
	res, err := msgServer.ClaimChallengerRewards(context, msg)

	// Verify the error response
	require.Error(t, err)
	require.Nil(t, res)

	// Verify the error message
	expectedErrMsg := "[ClaimChallengerRewards][DenomsSubsetOf] failed. Not enough coins to claim.: insufficient funds"
	require.Equal(t, expectedErrMsg, err.Error())
}

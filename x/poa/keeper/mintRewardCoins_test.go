package keeper_test

import (
	"soarchain/x/epoch/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMintRewardCoins(t *testing.T) {
	_, k, context, ctrl, bankMock, epochMock := SetupMsgServerClaimMotusRewards(t)
	// Set up the bank expectations
	bankMock.ExpectAny(context)
	epochMock.ExpectAny(context)
	defer ctrl.Finish()

	// Set up the context
	ctx := sdk.UnwrapSDKContext(context)

	//Set up some example epoch data
	epochData := types.EpochData{
		TotalEpochs:                   1,
		EpochV2VRX:                    "0udmotus",
		EpochV2VBX:                    "0udmotus",
		EpochV2NBX:                    "0udmotus",
		EpochRunner:                   "0udmotus",
		EpochChallenger:               "0udmotus",
		V2VRXtotalChallenges:          0,
		V2VBXtotalChallenges:          0,
		V2NBXtotalChallenges:          0,
		ChallengerTotalChallenges:     0,
		RunnerTotalChallenges:         0,
		TotalChallengesPrevDay:        0,
		InitialPerChallengeValue:      12.0,
		V2VRXLastBlockChallenges:      100,
		V2VRXPerChallengeValue:        2,
		V2VBXLastBlockChallenges:      150,
		V2VBXPerChallengeValue:        3,
		V2NBXLastBlockChallenges:      200,
		V2NBXPerChallengeValue:        4,
		RunnerLastBlockChallenges:     250,
		RunnerPerChallengeValue:       5,
		ChallengerLastBlockChallenges: 300,
		ChallengerPerChallengeValue:   6,
	}

	// Set the epoch data in the keeper
	epochMock.SetEpochData(ctx, epochData)

	// Call the MintRewardCoins function
	k.MintRewardCoins(ctx)

	//Verify that the epoch data values are reset to 0
	// updatedEpochData, found := epochMock.GetEpochData(ctx)
	// require.True(t, found)
	// require.Zero(t, updatedEpochData.V2VRXLastBlockChallenges)
	// require.Zero(t, updatedEpochData.V2VBXLastBlockChallenges)
	// require.Zero(t, updatedEpochData.V2NBXLastBlockChallenges)
	// require.Zero(t, updatedEpochData.RunnerLastBlockChallenges)
	// require.Zero(t, updatedEpochData.ChallengerLastBlockChallenges)

}

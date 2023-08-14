package keeper_test

import (
	"soarchain/x/epoch/types"
	"testing"
)

func TestMintRewardCoins(t *testing.T) {
	_, k, context, ctrl, _, epochMock := SetupMsgServerClaimMotusRewards(t)
	// Set up the bank expectations
	epochMock.ExpectAny(context)

	// // Execute the function to be tested
	// keeper.MintRewardCoins(ctx)

	//Set up some example epoch data
	epochData := types.EpochData{
		TotalEpochs:                   1,
		EpochV2VRX:                    "0udmotus",
		EpochV2VBX:                    "0udmotus",
		EpochV2NBX:                    "0udmotus",
		EpochRunner:                   "0udmotus",
		EpochChallenger:               "0udmotus",
		V2VRXTotalChallenges:          0,
		V2VBXTotalChallenges:          0,
		V2NBXTotalChallenges:          0,
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

	ctrl.Finish()

}

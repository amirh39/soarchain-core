package keeper_test

import (
	"soarchain/x/epoch/keeper"
	"soarchain/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateEpochData(keeper *keeper.Keeper, ctx sdk.Context) types.EpochData {
	item := types.EpochData{
		TotalEpochs:                   30,
		EpochV2VRX:                    "2udmotus",
		EpochV2VBX:                    "3udmotus",
		EpochV2NBX:                    "4udmotus",
		EpochRunner:                   "5udmotus",
		EpochChallenger:               "6",
		V2VRXTotalChallenges:          7,
		V2VBXTotalChallenges:          8,
		V2NBXTotalChallenges:          9,
		RunnerTotalChallenges:         10,
		ChallengerTotalChallenges:     11,
		V2VRXLastBlockChallenges:      1,
		V2VBXLastBlockChallenges:      1,
		V2NBXLastBlockChallenges:      1,
		RunnerLastBlockChallenges:     1,
		ChallengerLastBlockChallenges: 1,
		ChallengerPerChallengeValue:   1000000,
		V2NBXPerChallengeValue:        3000000,
		RunnerPerChallengeValue:       1000000,
		InitialPerChallengeValue:      9000000.0,
		TotalChallengesPrevDay:        99,
		V2VBXPerChallengeValue:        3000000,
		V2VRXPerChallengeValue:        3000000,
	}
	keeper.SetEpochData(ctx, item)
	return item
}

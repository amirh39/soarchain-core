package keeper_test

import (
	"soarchain/x/epoch/keeper"
	"soarchain/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateEpochData(keeper *keeper.Keeper, ctx sdk.Context) types.EpochData {
	item := types.EpochData{
		TotalEpochs:               30,
		EpochV2VRX:                "2utmotus",
		EpochV2VBX:                "3utmotus",
		EpochV2NBX:                "4utmotus",
		EpochRunner:               "5utmotus",
		EpochChallenger:           "6",
		V2VRXTotalChallenges:      7,
		V2VBXTotalChallenges:      8,
		V2NBXTotalChallenges:      9,
		RunnerTotalChallenges:     10,
		ChallengerTotalChallenges: 11,
	}
	keeper.SetEpochData(ctx, item)
	return item
}

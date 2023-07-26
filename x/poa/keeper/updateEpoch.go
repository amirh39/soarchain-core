package keeper

import (
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateEpoch updates the epoch data by calculating the total challenges for the previous day also it resets the epoch.
func (k Keeper) UpdateEpoch(ctx sdk.Context) {
	// Get the current epoch data
	epochData, _ := k.GetEpochData(ctx)

	// Calculate the total challenges for the previous day
	totalChallengesPrevDay := calculateTotalChallenges(epochData)

	// Increment the epoch count
	newEpochCnt := epochData.TotalEpochs + 1

	// Create the new epoch data with updated total challenges
	newEpochData := types.EpochData{
		TotalEpochs:                   newEpochCnt,
		EpochV2VRX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochV2VBX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochV2NBX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochRunner:                   sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochChallenger:               sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		V2VRXtotalChallenges:          0,
		V2VBXtotalChallenges:          0,
		V2NBXtotalChallenges:          0,
		RunnerTotalChallenges:         0,
		ChallengerTotalChallenges:     0,
		V2VRXLastBlockChallenges:      0,
		V2VBXLastBlockChallenges:      0,
		V2NBXLastBlockChallenges:      0,
		RunnerLastBlockChallenges:     0,
		ChallengerLastBlockChallenges: 0,
		TotalChallengesPrevDay:        totalChallengesPrevDay,
		InitialPerChallengeValue:      epochData.InitialPerChallengeValue,
		V2NBXPerChallengeValue:        epochData.V2NBXPerChallengeValue,
		RunnerPerChallengeValue:       epochData.RunnerPerChallengeValue,
		ChallengerPerChallengeValue:   epochData.ChallengerPerChallengeValue,
		V2VBXPerChallengeValue:        epochData.V2VBXPerChallengeValue,
		V2VRXPerChallengeValue:        epochData.V2VRXPerChallengeValue,
	}

	// Set the updated epoch data
	k.SetEpochData(ctx, newEpochData)
}

// calculateTotalChallenges calculates the total challenges for the previous day.
func calculateTotalChallenges(epochData types.EpochData) uint64 {
	epochData.TotalChallengesPrevDay += epochData.ChallengerTotalChallenges
	return epochData.TotalChallengesPrevDay
}

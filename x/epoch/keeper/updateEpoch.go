package keeper

import (
	params "github.com/soar-robotics/soarchain-core/app/params"
	"github.com/soar-robotics/soarchain-core/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateEpoch updates the epoch data by calculating the total challenges for the previous day also it resets the epoch.
func (k Keeper) UpdateEpoch(ctx sdk.Context) {
	// Get the current epoch data
	epochData, _ := k.GetEpochData(ctx)

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
		V2VRXTotalChallenges:          0,
		V2VBXTotalChallenges:          0,
		V2NBXTotalChallenges:          0,
		RunnerTotalChallenges:         0,
		ChallengerTotalChallenges:     0,
		V2VRXLastBlockChallenges:      epochData.V2VRXLastBlockChallenges,
		V2VBXLastBlockChallenges:      epochData.V2VBXLastBlockChallenges,
		V2NBXLastBlockChallenges:      epochData.V2NBXLastBlockChallenges,
		RunnerLastBlockChallenges:     epochData.RunnerLastBlockChallenges,
		ChallengerLastBlockChallenges: epochData.ChallengerLastBlockChallenges,
		TotalChallengesPrevDay:        epochData.TotalChallengesPrevDay,
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

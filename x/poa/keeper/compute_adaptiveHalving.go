package keeper

import (
	"log"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func divideMintedPerChallenge(mintedPerChallenge float64) (runner, challenger, v2nbx int64) {
	// Calculate 10% of mintedPerChallenge for both runner and challenger
	runner = int64(mintedPerChallenge * 0.10)
	challenger = runner // challenger gets the same percentage as runner

	// Calculate 80% of mintedPerChallenge for v2nbx
	v2nbx = int64(mintedPerChallenge * 0.80)

	return runner, challenger, v2nbx
}

func (k Keeper) ComputeAdaptiveHalving(ctx sdk.Context) error {
	epochData, isFound := k.GetEpochData(ctx)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[computeAdaptiveHalving][GetEpochData] failed. Epoch data is not found!")
	}

	targetValue := 600000.0
	totalChallengesTarget1 := 250_000_000

	A, B, C := utility.CalculateCoefficients(float64(epochData.InitialPerChallengeValue), targetValue, totalChallengesTarget1)
	log.Println("A B C = ", A, B, C)
	mintedPerChallenge, err := utility.CalculateMintedPerChallenge(epochData.InitialPerChallengeValue, int(epochData.TotalChallengesPrevDay), A, B, C)
	if err != nil {
		return sdkerrors.Wrap(err, "[computeAdaptiveHalving] failed to calculate minted per challenge")
	}

	log.Println("mintedPerChallenge= ", mintedPerChallenge)

	runner, challenger, v2nbx := divideMintedPerChallenge(mintedPerChallenge)

	// Update the initial per challenge value in epochData
	newEpochData := types.EpochData{
		TotalEpochs:                   epochData.TotalEpochs,
		EpochV2VRX:                    epochData.EpochV2VRX,
		EpochV2VBX:                    epochData.EpochV2VBX,
		EpochV2NBX:                    epochData.EpochV2NBX,
		EpochRunner:                   epochData.EpochRunner,
		EpochChallenger:               epochData.EpochChallenger,
		V2VRXtotalChallenges:          epochData.V2VRXtotalChallenges,
		V2VBXtotalChallenges:          epochData.V2VBXtotalChallenges,
		V2NBXtotalChallenges:          epochData.V2NBXtotalChallenges,
		RunnerTotalChallenges:         epochData.RunnerTotalChallenges,
		ChallengerTotalChallenges:     epochData.ChallengerTotalChallenges,
		V2VRXLastBlockChallenges:      epochData.V2VRXLastBlockChallenges,
		V2VBXLastBlockChallenges:      epochData.V2VBXLastBlockChallenges,
		V2NBXLastBlockChallenges:      epochData.V2NBXLastBlockChallenges,
		RunnerLastBlockChallenges:     epochData.RunnerLastBlockChallenges,
		ChallengerLastBlockChallenges: epochData.ChallengerLastBlockChallenges,
		TotalChallengesPrevDay:        0,
		InitialPerChallengeValue:      mintedPerChallenge,
		V2NBXPerChallengeValue:        uint64(v2nbx),
		RunnerPerChallengeValue:       uint64(runner),
		ChallengerPerChallengeValue:   uint64(challenger),
		V2VBXPerChallengeValue:        2000000,
		V2VRXPerChallengeValue:        2000000,
	}
	k.SetEpochData(ctx, newEpochData)

	return nil
}

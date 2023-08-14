package keeper

import (
	"log"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility"

	epoch "soarchain/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func divideMintedPerChallenge(mintedPerChallenge float64) (runner, challenger, v2nbx int64, err error) {
	// Check if mintedPerChallenge is a non-negative value
	if mintedPerChallenge < 0 {
		return 0, 0, 0, sdkerrors.Wrap(err, "mintedPerChallenge cannot be negative")
	}

	// Calculate 10% of mintedPerChallenge for both runner and challenger
	runner = int64(mintedPerChallenge * 0.10)
	challenger = runner // challenger gets the same percentage as runner

	// Calculate 80% of mintedPerChallenge for v2nbx
	v2nbx = int64(mintedPerChallenge * 0.80)

	return runner, challenger, v2nbx, nil
}

func (k Keeper) ComputeAdaptiveHalving(ctx sdk.Context) error {
	epochData, isFound := k.epochKeeper.GetEpochData(ctx)
	logger := k.Logger(ctx)
	log.Println("############## ComputeAdaptiveHalving Has Started ##############")
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[computeAdaptiveHalving] [GetEpochData] failed. Epoch data is not found!")
	}

	A, B, C, err := utility.CalculateCoefficients(float64(epochData.InitialPerChallengeValue), constants.TargetValue, constants.TotalChallengesTarget1)
	if err != nil {
		return sdkerrors.Wrap(err, "[ComputeAdaptiveHalving] [CalculateCoefficients] failed to calculate coefficients")
	}
	if logger != nil {
		logger.Info(" CalculateCoefficients successfully done.", A, B, C)
	}

	mintedPerChallenge, err := utility.CalculateMintedPerChallenge(epochData.InitialPerChallengeValue, int(epochData.TotalChallengesPrevDay), A, B, C)
	if err != nil {
		return sdkerrors.Wrap(err, "[computeAdaptiveHalving] [CalculateMintedPerChallenge] failed to calculate minted per challenge.")
	}
	if logger != nil {
		logger.Info(" Updated mintedPerChallenge value for today = ", mintedPerChallenge)
	}

	runner, challenger, v2nbx, err := divideMintedPerChallenge(mintedPerChallenge)
	if err != nil {
		return sdkerrors.Wrap(err, "[computeAdaptiveHalving] failed to calculate minted per challenge for objects")
	}

	if logger != nil {
		logger.Info("divideMintedPerChallenge successfully done.")
	}

	//Update the initial per challenge value in epochData
	newEpochData := epoch.EpochData{
		TotalEpochs:                   epochData.TotalEpochs,
		EpochV2VRX:                    epochData.EpochV2VRX,
		EpochV2VBX:                    epochData.EpochV2VBX,
		EpochV2NBX:                    epochData.EpochV2NBX,
		EpochRunner:                   epochData.EpochRunner,
		EpochChallenger:               epochData.EpochChallenger,
		V2VRXTotalChallenges:          epochData.V2VRXTotalChallenges,
		V2VBXTotalChallenges:          epochData.V2VBXTotalChallenges,
		V2NBXTotalChallenges:          epochData.V2NBXTotalChallenges,
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
	k.epochKeeper.SetEpochData(ctx, newEpochData)

	log.Println("############## End of ComputeAdaptiveHalving ##############")

	return nil
}

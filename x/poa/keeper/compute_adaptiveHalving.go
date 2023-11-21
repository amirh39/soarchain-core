package keeper

import (
	"log"

	"github.com/soar-robotics/soarchain-core/x/poa/constants"
	"github.com/soar-robotics/soarchain-core/x/poa/utility"

	epoch "github.com/soar-robotics/soarchain-core/x/epoch/types"

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

func (k Keeper) ComputeAdaptiveHalving(ctx sdk.Context, epoch epoch.EpochData) (epoch.EpochData, error) {
	logger := k.Logger(ctx)

	A, B, C, err := utility.CalculateCoefficients(float64(epoch.InitialPerChallengeValue), constants.TargetValue, constants.TotalChallengesTarget1)
	if err != nil {
		return epoch, sdkerrors.Wrap(err, "[ComputeAdaptiveHalving][CalculateCoefficients] failed to calculate coefficients")
	}
	if logger != nil {
		logger.Info(" CalculateCoefficients successfully done.", A, B, C)
	}

	mintedPerChallenge, err := utility.CalculateMintedPerChallenge(epoch.InitialPerChallengeValue, int(epoch.TotalChallengesPrevDay), A, B, C)
	if err != nil {
		return epoch, sdkerrors.Wrap(err, "[computeAdaptiveHalving][CalculateMintedPerChallenge] failed to calculate minted per challenge.")
	}
	if logger != nil {
		logger.Info(" Updated mintedPerChallenge value for today = ", mintedPerChallenge)
	}

	runner, challenger, v2nbx, err := divideMintedPerChallenge(mintedPerChallenge)
	if err != nil {
		return epoch, sdkerrors.Wrap(err, "[computeAdaptiveHalving][divideMintedPerChallenge] failed to calculate minted per challenge for objects")
	}

	if logger != nil {
		logger.Info("divideMintedPerChallenge successfully done.")
	}

	// Update the fields of epoch with the new values
	epoch.InitialPerChallengeValue = mintedPerChallenge
	epoch.V2NBXPerChallengeValue = uint64(v2nbx)
	epoch.RunnerPerChallengeValue = uint64(runner)
	epoch.ChallengerPerChallengeValue = uint64(challenger)
	epoch.V2VBXPerChallengeValue = 2000000
	epoch.V2VRXPerChallengeValue = 2000000
	epoch.TotalChallengesPrevDay = 0

	log.Println("############## End of ComputeAdaptiveHalving ##############")

	return epoch, nil
}

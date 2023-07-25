package keeper

import (
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ComputeAdaptiveHalving(ctx sdk.Context) error {
	epochData, isFound := k.GetEpochData(ctx)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[computeAdaptiveHalving][GetEpochData] failed. Epoch data is not found!")
	}

	targetValue := 600000.0
	totalChallengesTarget1 := 250_000_000

	A, B, C := utility.CalculateCoefficients(float64(epochData.InitialPerChallengeValue), targetValue, totalChallengesTarget1)

	mintedPerChallenge, err := utility.CalculateMintedPerChallenge(float64(epochData.GetInitialPerChallengeValue()), int(epochData.GetTotalChallengesPrevDay()), A, B, C)
	if err != nil {
		return sdkerrors.Wrap(err, "[computeAdaptiveHalving] failed to calculate minted per challenge")
	}

	// Update the initial per challenge value in epochData
	epochData.InitialPerChallengeValue = uint64(mintedPerChallenge)

	return nil
}

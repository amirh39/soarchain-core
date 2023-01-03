package keeper

import (
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) V2NRewardCalculator(ctx sdk.Context, rewardMultiplier float64, clientCommunicationMode string) (float64, error) {

	rewardPerBlock, err := utility.V2NRewardEmissionPerBlock(ctx, clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "V2V Motus Reward Emission per block couldn't be computed. Check client communication mode.")
	}

	// Score is below 50, no rewards are earned
	if rewardMultiplier == 0 {
		return 0, nil
	}

	var totalMultipliers float64 = 0.0

	if clientCommunicationMode == "v2n-bx" {
		allClients := k.GetAllClient(ctx)

		for i := 0; i < len(allClients); i++ {
			currMultiplier, err := strconv.ParseFloat(allClients[i].RewardMultiplier, 64)
			if err != nil {
				return 0.0, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
			}
			totalMultipliers += currMultiplier
		}
	} else if clientCommunicationMode == "runner" {
		allRunners := k.GetAllRunner(ctx)

		for i := 0; i < len(allRunners); i++ {
			currMultiplier, err := strconv.ParseFloat(allRunners[i].RewardMultiplier, 64)
			if err != nil {
				return 0.0, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
			}
			totalMultipliers += currMultiplier
		}
	} else {
		return 0.0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "V2N communication mode is not supported!")
	}

	// Protection against +Inf netEarnings calculation
	if totalMultipliers > 0 {
		return (rewardMultiplier / totalMultipliers) * rewardPerBlock, nil
	} else {
		return 0, nil
	}

}

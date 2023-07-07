package keeper

import (
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) V2NRewardCalculator(ctx sdk.Context, rewardMultiplier float64, clientCommunicationMode string) (float64, error) {
	logger := k.Logger(ctx)

	rewardPerBlock, err := utility.V2NRewardEmissionPerBlock(ctx, clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[V2NRewardCalculator][V2NRewardEmissionPerBlock] failed. V2V Motus Reward Emission per block couldn't be computed. Check client communication mode. Error: [ %T ]", err)
	}

	if logger != nil {
		logger.Info("V2V Motus Reward Emission per block successfully computed.", "calculation of", "V2NRewardCalculator")
	}

	// Score is below 50, no rewards are earned
	if rewardMultiplier == 0 || rewardMultiplier < 0 {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[V2NRewardCalculator] failed. Reward Multiplier can not be negative. got [ %T ]. Error: [ %T ]", rewardMultiplier, err)
	}

	var totalMultipliers float64 = 0.0

	if clientCommunicationMode == constants.V2NBX {
		allClients := k.GetAllClient(ctx)
		if allClients == nil {
			return 0, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[V2NRewardCalculator][GetAllClient] failed. Couldn't find any client for v2n-bx. Error: [ %T ]", err)
		}

		for i := 0; i < len(allClients); i++ {
			currMultiplier, err := strconv.ParseFloat(allClients[i].RewardMultiplier, 64)
			if err != nil {
				return 0.0, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[V2NRewardCalculator][ParseFloat] failed. Couldn't convert the string to a floating-point number for a v2n-bx. Error: [ %T ]", err)
			}
			totalMultipliers += currMultiplier
		}
		if logger != nil {
			logger.Info("V2NBX total multiplier successfully computed.", "calculation of", "V2NRewardCalculator")
		}
	} else if clientCommunicationMode == constants.Runner {
		allRunners := k.GetAllRunner(ctx)
		if allRunners == nil {
			return 0, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[V2NRewardCalculator][GetAllRunner] failed. Couldn't find any client for runner. Error: [ %T ]", err)
		}

		for i := 0; i < len(allRunners); i++ {
			currMultiplier, err := strconv.ParseFloat(allRunners[i].RewardMultiplier, 64)
			if err != nil {
				return 0.0, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[V2NRewardCalculator][ParseFloat] failed. Couldn't convert the string to a floating-point number for a runner. Error: [ %T ]", err)
			}
			totalMultipliers += currMultiplier
		}
		if logger != nil {
			logger.Info("Runner total multiplier successfully computed.", "calculation of", "V2NRewardCalculator")
		}
	} else {
		return 0.0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2NRewardCalculator][ParseFloat] failed. V2N communication mode is not supported.")
	}

	// Protection against +Inf netEarnings calculation
	if totalMultipliers > 0 {
		return (rewardMultiplier / totalMultipliers) * float64(rewardPerBlock), nil
	} else {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrLogic, "[V2NRewardCalculator] failed. Sum of Multipliers can not be negative OR zero. Error: [ %T ]", err)
	}

}

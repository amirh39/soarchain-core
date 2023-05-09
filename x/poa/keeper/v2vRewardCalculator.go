package keeper

import (
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) V2VRewardCalculator(ctx sdk.Context, rewardMultiplier float64, clientCommunicationMode string) (float64, error) {

	rewardPerBlock, err := utility.V2VRewardEmissionPerBlock(ctx, clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrPanic, "[V2VRewardCalculator][V2VRewardEmissionPerBlock] failed. V2V Motus Reward Emission per block couldn't be calculated. Check client communication mode. Error: [ %T ]", err)
	}

	// Score is below 50, no rewards are earned
	if rewardMultiplier == 0 || rewardMultiplier < 0 {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[V2VRewardCalculator] failed. Reward Multiplier can not be negative. got [ %T ]. Error: [ %T ]", rewardMultiplier, err)
	}

	allClients := k.GetAllClient(ctx)
	if allClients == nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[V2VRewardCalculator][GetAllClient] failed. Couldn't find any client for v2n-bx. Error: [ %T ]", err)
	}

	var totalMultipliers float64 = 0.0

	for i := 0; i < len(allClients); i++ {
		currMultiplier, err := strconv.ParseFloat(allClients[i].RewardMultiplier, 64)
		if err != nil {
			return 0.0, sdkerrors.Wrapf(sdkerrors.ErrPanic, "[V2VRewardCalculator][ParseFloat] failed. Couldn't convert the string to a floating-point number. Error: [ %T ]", err)
		}
		totalMultipliers += currMultiplier
	}

	// Protection against +Inf netEarnings calculation
	if totalMultipliers > 0 {
		return (rewardMultiplier / totalMultipliers) * float64(rewardPerBlock), nil
	} else {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrLogic, "[V2VRewardCalculator] failed. Sum of Multipliers can not be negative OR zero. Error: [ %T ]", err)
	}
}

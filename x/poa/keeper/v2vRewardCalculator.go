package keeper

import (
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) V2VRewardCalculator(ctx sdk.Context, rewardMultiplier float64, clientCommunicationMode string) (float64, error) {

	rewardPerBlock, err := utility.V2VRewardEmissionPerBlock(ctx, clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "[V2VRewardCalculator][V2VRewardEmissionPerBlock] failed. V2V Motus Reward Emission per block couldn't be calculated. Check client communication mode."+err.Error())
	}

	// Score is below 50, no rewards are earned
	if rewardMultiplier == 0 {
		return 0, nil
	}

	allClients := k.GetAllClient(ctx)
	var totalMultipliers float64 = 0.0

	for i := 0; i < len(allClients); i++ {
		currMultiplier, err := strconv.ParseFloat(allClients[i].RewardMultiplier, 64)
		if err != nil {
			return 0.0, sdkerrors.Wrap(sdkerrors.ErrPanic, "[V2VRewardCalculator][V2VRewardEmissionPerBlock] failed. Couldn't convert and parse all client reward to Float64."+err.Error())
		}
		totalMultipliers += currMultiplier
	}

	// Protection against +Inf netEarnings calculation
	if totalMultipliers > 0 {
		return (rewardMultiplier / totalMultipliers) * float64(rewardPerBlock), nil
	} else {
		return 0, nil
	}

}

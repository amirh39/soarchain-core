package keeper

import (
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MotusReward(ctx sdk.Context, rewardMultiplier float64) (float64, error) {

	allClients := k.GetAllClient(ctx)
	var totalMultipliers float64 = 0.0

	for i := 0; i < len(allClients); i++ {
		currMultiplier, err := strconv.ParseFloat(allClients[i].RewardMultiplier, 64)
		if err != nil {
			return 0.0, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		totalMultipliers += currMultiplier
	}

	rewardPerBlock := utility.V2VReceiveRewardEmissionPerBlock(ctx)

	return (rewardMultiplier / totalMultipliers) * rewardPerBlock, nil

}

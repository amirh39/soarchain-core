package keeper

import (
	param "soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateReward(totalAmount, score float64) sdk.Coin {
	// Calculate the reward amount based on the total amount of tokens and score
	reward := (totalAmount / 100.0) * score

	earnedRewardsInt := sdk.NewIntFromUint64((uint64(reward)))
	earnedCoin := sdk.NewCoin(param.BondDenom, earnedRewardsInt)

	return earnedCoin
}

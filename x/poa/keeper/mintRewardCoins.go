package keeper

import (
	"strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {

	var client types.Client
	var totalV2VRewards int = 0
	clients := k.GetAllClient(ctx)
	for i := 0; i < len(clients); i++ {
		client = clients[i]
		clientReward, _ := strconv.Atoi(client.NetEarnings)
		totalV2VRewards += clientReward
	}

	totalV2VRewardsInt := sdk.NewIntFromUint64((uint64(totalV2VRewards)))
	coin := sdk.NewCoin("soar", totalV2VRewardsInt)
	epochSoarRewardCoins := sdk.Coins{coin}

	k.bankKeeper.MintCoins(ctx, types.ModuleName, epochSoarRewardCoins)

}

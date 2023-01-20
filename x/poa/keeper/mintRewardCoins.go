package keeper

import (
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {

	var client types.Client
	var totalV2VRewards sdk.Coin = sdk.NewCoin("soar", sdk.ZeroInt())

	clients := k.GetAllClient(ctx)
	if len(clients) > 0 {
		for i := 0; i < len(clients); i++ {
			client = clients[i]
			clientReward, _ := sdk.ParseCoinNormalized(client.NetEarnings)
			totalV2VRewards = totalV2VRewards.Add(clientReward)
		}

		epochSoarRewardCoins := sdk.Coins{totalV2VRewards}

		k.bankKeeper.MintCoins(ctx, types.ModuleName, epochSoarRewardCoins)
	}

}

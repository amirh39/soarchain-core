package keeper

import (
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {

	epochData, _ := k.GetEpochData(ctx)

	// Parse coins
	v2vRxRewards, _ := sdk.ParseCoinsNormalized(epochData.EpochV2VRX)
	v2vBxRewards, _ := sdk.ParseCoinsNormalized(epochData.EpochV2VBX)
	v2nBxRewards, _ := sdk.ParseCoinsNormalized(epochData.EpochV2NBX)
	runnerRewards, _ := sdk.ParseCoinsNormalized(epochData.EpochRunner)

	k.bankKeeper.MintCoins(ctx, types.ModuleName, v2vRxRewards)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, v2vBxRewards)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, v2nBxRewards)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, runnerRewards)

	// Reset Epoch Rewards
	newEpochData := types.EpochData{
		TotalEpochs: epochData.TotalEpochs,
		EpochV2VRX:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		EpochV2VBX:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		EpochV2NBX:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		EpochRunner: sdk.NewCoin("soar", sdk.ZeroInt()).String(),
	}
	k.SetEpochData(ctx, newEpochData)

	// var client types.Client
	// var totalV2VRewards sdk.Coin = sdk.NewCoin("soar", sdk.ZeroInt())

	// clients := k.GetAllClient(ctx)
	// if len(clients) > 0 {
	// 	for i := 0; i < len(clients); i++ {
	// 		client = clients[i]
	// 		clientReward, _ := sdk.ParseCoinNormalized(client.NetEarnings)
	// 		totalV2VRewards = totalV2VRewards.Add(clientReward)
	// 	}

	// 	epochSoarRewardCoins := sdk.Coins{totalV2VRewards}

	// 	k.bankKeeper.MintCoins(ctx, types.ModuleName, epochSoarRewardCoins)
	// }

}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SendLefoverCoins(ctx sdk.Context) {

	// var leftOverV2VRx sdk.Coin
	// var leftOverV2VBx sdk.Coin
	// var leftOverV2NBx sdk.Coin
	// var leftOverRunner sdk.Coin

	// // v2vrx
	// if v2vRxReward.IsLT(TargetV2VBxCoin) || v2vRxReward.IsEqual(targetV2VRxCoin) {
	// 	leftOverV2VRx = targetV2VRxCoin.Sub(v2vRxReward)
	// 	if leftOverV2VRx.IsPositive() {
	// 		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverV2VRx})
	// 	}
	// }

	// // v2vbx
	// if v2vBxReward.IsLT(targetV2VBxCoin) || v2vBxReward.IsEqual(targetV2VBxCoin) {
	// 	leftOverV2VBx = targetV2VBxCoin.Sub(v2vBxReward)
	// 	if leftOverV2VBx.IsPositive() {
	// 		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverV2VBx})
	// 	}
	// }

	// // v2nbx
	// if v2nBxReward.IsLT(targetV2NBxCoin) || v2nBxReward.IsEqual(targetV2NBxCoin) {
	// 	leftOverV2NBx = targetV2NBxCoin.Sub(v2nBxReward)
	// 	if leftOverV2NBx.IsPositive() {
	// 		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverV2NBx})
	// 	}
	// }

	// // runner
	// if runnerReward.IsLT(targetRunnerCoin) || runnerReward.IsEqual(targetRunnerCoin) {
	// 	leftOverRunner = targetRunnerCoin.Sub(runnerReward)
	// 	if leftOverRunner.IsPositive() {
	// 		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverRunner})
	// 	}
	// }
}

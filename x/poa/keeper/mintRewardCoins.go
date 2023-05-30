package keeper

import (
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	params "soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {

	// Calculate rewards
	targetV2VRx, _ := utility.V2VRewardEmissionPerEpoch(ctx, "v2v-rx")
	targetV2VRxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetV2VRx)))

	targetV2VBx, _ := utility.V2VRewardEmissionPerEpoch(ctx, "v2v-bx")
	targetV2VBxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetV2VBx)))

	targetV2NBx, _ := utility.V2NRewardEmissionPerEpoch(ctx, "v2n-bx")
	targetV2NBxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetV2NBx)))

	targetRunner, _ := utility.V2NRewardEmissionPerEpoch(ctx, "runner")
	targetRunnerCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetRunner)))

	v2vRxReward, _ := sdk.ParseCoinNormalized(targetV2VRxCoin.Denom)
	v2vBxReward, _ := sdk.ParseCoinNormalized(targetV2VBxCoin.Denom)
	v2nBxReward, _ := sdk.ParseCoinNormalized(targetV2NBxCoin.Denom)
	runnerReward, _ := sdk.ParseCoinNormalized(targetRunnerCoin.Denom)

	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{v2vRxReward})
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{v2vBxReward})
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{v2nBxReward})
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{runnerReward})

	epochData, _ := k.GetEpochData(ctx)
	// Reset Epoch Rewards
	newEpochData := types.EpochData{

		TotalEpochs: epochData.TotalEpochs,
		EpochV2VRX:  v2vRxReward.String(),
		EpochV2VBX:  v2vBxReward.String(),
		EpochV2NBX:  v2nBxReward.String(),
		EpochRunner: runnerReward.String(),
	}
	k.SetEpochData(ctx, newEpochData)

	var leftOverV2VRx sdk.Coin
	var leftOverV2VBx sdk.Coin
	var leftOverV2NBx sdk.Coin
	var leftOverRunner sdk.Coin

	// v2vrx
	if v2vRxReward.IsLT(targetV2VRxCoin) {
		leftOverV2VRx = targetV2VRxCoin.Sub(v2vRxReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverV2VRx})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverV2VRx})
	}
	// v2vbx
	if v2vBxReward.IsLT(targetV2VBxCoin) {
		leftOverV2VBx = targetV2VBxCoin.Sub(v2vBxReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverV2VBx})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverV2VBx})
	}
	//v2nbx
	if v2nBxReward.IsLT(targetV2NBxCoin) {
		leftOverV2NBx = targetV2NBxCoin.Sub(v2nBxReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverV2NBx})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverV2NBx})
	}
	// runner
	if runnerReward.IsLT(targetRunnerCoin) {
		leftOverRunner = targetRunnerCoin.Sub(runnerReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverRunner})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "rewardcap", sdk.Coins{leftOverRunner})
	}

}

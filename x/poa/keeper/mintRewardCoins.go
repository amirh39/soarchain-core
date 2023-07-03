package keeper

import (
	"log"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	params "soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {
	logger := k.Logger(ctx)
	epochData, _ := k.GetEpochData(ctx)

	log.Println("############## Minting Reward Coins Transaction Started ##############")

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
		EpochV2VRX:  sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochV2VBX:  sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochV2NBX:  sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochRunner: sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
	}
	k.SetEpochData(ctx, newEpochData)

	if logger != nil {
		logger.Info("Epoch Rewards successfully updated.", "transaction", "MintRewardCoins")
	}

	// Calculate leftover rewards
	targetV2VRx, _ := utility.V2VRewardEmissionPerEpoch(ctx, constants.V2VRX)
	targetV2VRxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetV2VRx)))

	targetV2VBx, _ := utility.V2VRewardEmissionPerEpoch(ctx, constants.V2VBX)
	targetV2VBxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetV2VBx)))

	targetV2NBx, _ := utility.V2NRewardEmissionPerEpoch(ctx, constants.V2NBX)
	targetV2NBxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetV2NBx)))

	targetRunner, _ := utility.V2NRewardEmissionPerEpoch(ctx, constants.Runner)
	targetRunnerCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetRunner)))

	if logger != nil {
		logger.Info("Leftover rewards successfully Calculated.", "transaction", "MintRewardCoins")
	}

	//
	v2vRxReward, _ := sdk.ParseCoinNormalized(epochData.EpochV2VRX)
	v2vBxReward, _ := sdk.ParseCoinNormalized(epochData.EpochV2VBX)
	v2nBxReward, _ := sdk.ParseCoinNormalized(epochData.EpochV2NBX)
	runnerReward, _ := sdk.ParseCoinNormalized(epochData.EpochRunner)

	var leftOverV2VRx sdk.Coin
	var leftOverV2VBx sdk.Coin
	var leftOverV2NBx sdk.Coin
	var leftOverRunner sdk.Coin

	// v2vrx
	if v2vRxReward.IsLT(targetV2VRxCoin) {
		leftOverV2VRx = targetV2VRxCoin.Sub(v2vRxReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverV2VRx})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, constants.RewardCap, sdk.Coins{leftOverV2VRx})
	}
	// v2vbx
	if v2vBxReward.IsLT(targetV2VBxCoin) {
		leftOverV2VBx = targetV2VBxCoin.Sub(v2vBxReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverV2VBx})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, constants.RewardCap, sdk.Coins{leftOverV2VBx})
	}
	//v2nbx
	if v2nBxReward.IsLT(targetV2NBxCoin) {
		leftOverV2NBx = targetV2NBxCoin.Sub(v2nBxReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverV2NBx})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, constants.RewardCap, sdk.Coins{leftOverV2NBx})
	}
	// runner
	if runnerReward.IsLT(targetRunnerCoin) {
		leftOverRunner = targetRunnerCoin.Sub(runnerReward)
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{leftOverRunner})
		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, constants.RewardCap, sdk.Coins{leftOverRunner})
	}
	if logger != nil {
		logger.Info("MintRewardCoins successfully submitted in the chain.", "transaction", "MintRewardCoins")
	}

	log.Println("############## End of  Minting Reward Coins Transaction ##############")
}

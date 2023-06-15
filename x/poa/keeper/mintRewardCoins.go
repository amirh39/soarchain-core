package keeper

import (
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	params "soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {

	// Calculate rewards
	TargetV2VRx, _ := utility.V2VRewardEmissionPerEpoch(ctx, "v2v-rx")
	TargetV2VRxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(TargetV2VRx)))

	TargetV2VBx, _ := utility.V2VRewardEmissionPerEpoch(ctx, "v2v-bx")
	TargetV2VBxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(TargetV2VBx)))

	TargetV2NBx, _ := utility.V2NRewardEmissionPerEpoch(ctx, "v2n-bx")
	TargetV2NBxCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(TargetV2NBx)))

	TargetRunner, _ := utility.V2NRewardEmissionPerEpoch(ctx, "runner")
	TargetRunnerCoin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(TargetRunner)))

	v2vRxReward, _ := sdk.ParseCoinNormalized(TargetV2VRxCoin.String())
	v2vBxReward, _ := sdk.ParseCoinNormalized(TargetV2VBxCoin.String())
	v2nBxReward, _ := sdk.ParseCoinNormalized(TargetV2NBxCoin.String())
	runnerReward, _ := sdk.ParseCoinNormalized(TargetRunnerCoin.String())

	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{v2vRxReward})
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{v2vBxReward})
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{v2nBxReward})
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{runnerReward})

	epochData, _ := k.GetEpochData(ctx)
	// Reset Epoch Rewards
	epochCnt := epochData.TotalEpochs
	newEpochCnt := epochCnt + 1

	newEpochData := types.EpochData{

		TotalEpochs: newEpochCnt,
		EpochV2VRX:  v2vRxReward.String(),
		EpochV2VBX:  v2vBxReward.String(),
		EpochV2NBX:  v2nBxReward.String(),
		EpochRunner: runnerReward.String(),
	}
	k.SetEpochData(ctx, newEpochData)

}

package keeper

import (
	"log"
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {
	epochData, _ := k.GetEpochData(ctx)

	if epochData.V2VRXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epochData.V2VRXLastBlockChallenges, int(epochData.V2VRXPerChallengeValue))
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.V2VRXLastBlockChallenges = 0
	}

	if epochData.V2VBXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epochData.V2VBXLastBlockChallenges, int(epochData.V2VBXPerChallengeValue))
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.V2VBXLastBlockChallenges = 0
	}

	if epochData.V2NBXLastBlockChallenges != 0 {
		log.Println(epochData.V2NBXLastBlockChallenges)
		log.Println(epochData.V2NBXPerChallengeValue)
		rewardToSet := epochData.V2NBXPerChallengeValue * epochData.V2NBXLastBlockChallenges
		log.Println(rewardToSet)
		coin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(rewardToSet))
		log.Println(coin)
		log.Println("V2NBXLastBlockChallenges=", epochData.V2NBXLastBlockChallenges)
		mintAndParseCoins(ctx, coin, k)
		epochData.V2NBXLastBlockChallenges = 0
	}

	if epochData.RunnerLastBlockChallenges != 0 {
		log.Println(epochData.RunnerLastBlockChallenges)
		log.Println(epochData.RunnerPerChallengeValue)
		rewardToSet := epochData.RunnerPerChallengeValue * uint64(epochData.RunnerLastBlockChallenges)
		log.Println(rewardToSet)
		coin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(rewardToSet))
		log.Println(coin)
		log.Println("RunnerLastBlockChallenges=", epochData.RunnerLastBlockChallenges)
		mintAndParseCoins(ctx, coin, k)
		epochData.RunnerLastBlockChallenges = 0
	}

	if epochData.ChallengerLastBlockChallenges != 0 {
		log.Println(epochData.ChallengerLastBlockChallenges)
		log.Println(epochData.ChallengerPerChallengeValue)
		rewardToSet := epochData.ChallengerPerChallengeValue * uint64(epochData.ChallengerLastBlockChallenges)
		log.Println(rewardToSet)
		coin := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(rewardToSet))
		log.Println(coin)
		log.Println("ChallengerLastBlockChallenges=", epochData.ChallengerLastBlockChallenges)
		mintAndParseCoins(ctx, coin, k)
		epochData.ChallengerLastBlockChallenges = 0
	}

	k.SetEpochData(ctx, epochData)
}
func parseUintAndCreateCoin(value uint64, multiplier int) sdk.Coin {
	amount := value * uint64(multiplier)
	return sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(amount))
}

func mintAndParseCoins(ctx sdk.Context, coin sdk.Coin, k Keeper) {
	parsedCoin, _ := sdk.ParseCoinNormalized(coin.String())
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{parsedCoin})
}

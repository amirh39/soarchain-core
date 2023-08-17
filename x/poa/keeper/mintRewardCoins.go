package keeper

import (
	params "soarchain/app/params"
	epoch "soarchain/x/epoch/types"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context, epoch epoch.EpochData) error {
	logger := k.Logger(ctx)

	if epoch.V2VRXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epoch.V2VRXLastBlockChallenges, int(epoch.V2VRXPerChallengeValue))
		if logger != nil {
			logger.Info("V2VRXLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epoch.V2VRXLastBlockChallenges = 0
	}

	if epoch.V2VBXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epoch.V2VBXLastBlockChallenges, int(epoch.V2VBXPerChallengeValue))
		if logger != nil {
			logger.Info("V2VBXLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epoch.V2VBXLastBlockChallenges = 0
	}

	if epoch.V2NBXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epoch.V2NBXPerChallengeValue, int(epoch.V2NBXLastBlockChallenges))
		if logger != nil {
			logger.Info("V2NBXLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epoch.V2NBXLastBlockChallenges = 0
	}

	if epoch.RunnerLastBlockChallenges != 0 {

		rewardToSet := parseUintAndCreateCoin(epoch.RunnerPerChallengeValue, int(epoch.RunnerLastBlockChallenges))

		if logger != nil {
			logger.Info("RunnerLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epoch.RunnerLastBlockChallenges = 0
	}

	if epoch.ChallengerLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epoch.ChallengerPerChallengeValue, int(epoch.ChallengerLastBlockChallenges))
		if logger != nil {
			logger.Info("ChallengerLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epoch.ChallengerLastBlockChallenges = 0
	}

	k.epochKeeper.SetEpochData(ctx, epoch)

	return nil
}
func parseUintAndCreateCoin(value uint64, multiplier int) sdk.Coin {
	amount := value * uint64(multiplier)
	return sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(amount))
}

func mintAndParseCoins(ctx sdk.Context, coin sdk.Coin, k Keeper) {
	parsedCoin, _ := sdk.ParseCoinNormalized(coin.String())
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{parsedCoin})
}

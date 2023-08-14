package keeper

import (
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) error {
	logger := k.Logger(ctx)
	epochData, found := k.epochKeeper.GetEpochData(ctx)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[MintRewardCoins][GetEpochData] failed. Epoch data is not found!")
	}

	if epochData.V2VRXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epochData.V2VRXLastBlockChallenges, int(epochData.V2VRXPerChallengeValue))
		if logger != nil {
			logger.Info("V2VRXLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.V2VRXLastBlockChallenges = 0
	}

	if epochData.V2VBXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epochData.V2VBXLastBlockChallenges, int(epochData.V2VBXPerChallengeValue))
		if logger != nil {
			logger.Info("V2VBXLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.V2VBXLastBlockChallenges = 0
	}

	if epochData.V2NBXLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epochData.V2NBXPerChallengeValue, int(epochData.V2NBXLastBlockChallenges))
		if logger != nil {
			logger.Info("V2NBXLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.V2NBXLastBlockChallenges = 0
	}

	if epochData.RunnerLastBlockChallenges != 0 {

		rewardToSet := parseUintAndCreateCoin(epochData.RunnerPerChallengeValue, int(epochData.RunnerLastBlockChallenges))

		if logger != nil {
			logger.Info("RunnerLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.RunnerLastBlockChallenges = 0
	}

	if epochData.ChallengerLastBlockChallenges != 0 {
		rewardToSet := parseUintAndCreateCoin(epochData.ChallengerPerChallengeValue, int(epochData.ChallengerLastBlockChallenges))
		if logger != nil {
			logger.Info("ChallengerLastBlockChallenges successfully minted.", "transaction", "MintRewardCoins", "Minted amount", rewardToSet)
		}
		mintAndParseCoins(ctx, rewardToSet, k)
		epochData.ChallengerLastBlockChallenges = 0
	}

	k.epochKeeper.SetEpochData(ctx, epochData)

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

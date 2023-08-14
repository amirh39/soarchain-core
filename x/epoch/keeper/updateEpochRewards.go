package keeper

import (
	"log"
	"soarchain/x/epoch/types"
	"soarchain/x/poa/constants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) updateEpochDataForClientType(
	ctx sdk.Context,
	clientType string,
	rewardToSet sdk.Coin,
	epochData types.EpochData,
) (types.EpochData, error) {
	logger := k.Logger(ctx)

	var err error
	switch clientType {
	case constants.V2VRX:
		epochData.EpochV2VRX, err = updateCoin(epochData.EpochV2VRX, rewardToSet)
	case constants.V2VBX:
		epochData.EpochV2VBX, err = updateCoin(epochData.EpochV2VBX, rewardToSet)
	case constants.V2NBX:
		epochData.EpochV2NBX, err = updateCoin(epochData.EpochV2NBX, rewardToSet)
	case constants.Runner:
		epochData.EpochRunner, err = updateCoin(epochData.EpochRunner, rewardToSet)
	case constants.Challenger:
		epochData.EpochChallenger, err = updateCoin(epochData.EpochChallenger, rewardToSet)
	case constants.V2NChallenge:
		epochData, err = updateChallenges(epochData)
	default:
		return epochData, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[UpdateEpochRewards] failed. Client type is not valid.")
	}

	if err != nil {
		return epochData, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[UpdateEpochRewards] failed. UpdateCoin function failed.")
	}

	if logger != nil {
		logger.Info("Updating epoch data for client type.", "transaction", "UpdateEpochRewards", "clientType", clientType)
	}

	return epochData, nil
}

func updateCoin(existingCoinStr string, rewardToAdd sdk.Coin) (string, error) {
	existingCoin, err := sdk.ParseCoinNormalized(existingCoinStr)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[UpdateEpochRewards][ParseCoinsNormalized] failed. Amount: [%s] couldn't be parsed. Error: %s", existingCoinStr, err)
	}

	newCoin := existingCoin.Add(rewardToAdd)
	return newCoin.String(), nil
}

func updateChallenges(epochData types.EpochData) (types.EpochData, error) {
	epochData.V2NBXTotalChallenges++
	epochData.RunnerTotalChallenges++
	epochData.ChallengerTotalChallenges++
	epochData.V2NBXLastBlockChallenges++
	epochData.RunnerLastBlockChallenges++
	epochData.ChallengerLastBlockChallenges++
	epochData.TotalChallengesPrevDay++
	return epochData, nil
}

func (k Keeper) UpdateEpochRewards(ctx sdk.Context, clientType string, rewardToSet sdk.Coin) error {
	logger := k.Logger(ctx)
	log.Println("############## Update Epoch Rewards Started ##############")

	epochData, isFound := k.GetEpochData(ctx)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateEpochRewards][GetEpochData] failed. Epoch data is not found!")
	}

	if logger != nil {
		logger.Info("Getting epoch data successfully done.", "transaction", "UpdateEpochRewards", "epochData", epochData, "isFound", isFound)
	}

	updatedEpochData, err := k.updateEpochDataForClientType(ctx, clientType, rewardToSet, epochData)
	if err != nil {
		return sdkerrors.Wrap(err, "[UpdateEpochRewards][updateEpochDataForClientType] failed. updateEpochDataForClientType function failed.")
	}

	k.SetEpochData(ctx, updatedEpochData)

	log.Println("############## End of Update Epoch Rewards ##############")

	return nil
}

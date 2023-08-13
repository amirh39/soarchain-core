package keeper

import (
	"log"
	"soarchain/x/epoch/types"
	"soarchain/x/poa/constants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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

	if logger != nil {
		logger.Info("Print out the client type.", "transaction", "UpdateEpochRewards", "clientType", clientType)
	}

	switch clientType {

	case constants.V2VRX:

		if logger != nil {
			logger.Info("updating v2v-rx epochValue.", "UpdateEpochRewards")
		}
		// Parse the current value into a sdk.Coin
		epochV2VRXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2VRX)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[UpdateEpochRewards][ParseCoinsNormalized] failed. Amount: [ %T ] couldn't be parsed. Error: [ %T ]", epochV2VRXCoin, err)

		}

		// Addition rewardToSet
		newEpochV2VRXCoin := epochV2VRXCoin.Add(rewardToSet)

		// Convert the result back to a string representation
		newEpochV2VRX := newEpochV2VRXCoin.String()

		// Create a new EpochData object with the updated value
		newEpochData := types.EpochData{
			TotalEpochs:                 epochData.TotalEpochs,
			EpochV2VRX:                  newEpochV2VRX,
			EpochV2VBX:                  epochData.EpochV2VBX,
			EpochV2NBX:                  epochData.EpochV2NBX,
			EpochRunner:                 epochData.EpochRunner,
			EpochChallenger:             epochData.EpochChallenger,
			V2VRXTotalChallenges:        epochData.V2VRXTotalChallenges,
			V2VBXTotalChallenges:        epochData.V2VBXTotalChallenges,
			V2NBXTotalChallenges:        epochData.V2NBXTotalChallenges,
			RunnerTotalChallenges:       epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges:   epochData.ChallengerTotalChallenges,
			ChallengerPerChallengeValue: epochData.ChallengerPerChallengeValue,
			V2NBXPerChallengeValue:      epochData.V2NBXPerChallengeValue,
			RunnerPerChallengeValue:     epochData.RunnerPerChallengeValue,
			InitialPerChallengeValue:    epochData.InitialPerChallengeValue,
			TotalChallengesPrevDay:      epochData.TotalChallengesPrevDay,
		}

		// Store the updated epoch data
		k.SetEpochData(ctx, newEpochData)

	case constants.V2VBX:
		if logger != nil {
			logger.Info("updating v2v-bx epochValue.", "UpdateEpochRewards")
		}
		epochV2VBXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2VBX)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[UpdateEpochRewards][ParseCoinsNormalized] failed. Amount: [ %T ] couldn't be parsed. Error: [ %T ]", epochV2VBXCoin, err)
		}
		newEpochV2VBXCoin := epochV2VBXCoin.Add(rewardToSet)
		newEpochV2VBX := newEpochV2VBXCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:                 epochData.TotalEpochs,
			EpochV2VRX:                  epochData.EpochV2VRX,
			EpochV2VBX:                  newEpochV2VBX,
			EpochV2NBX:                  epochData.EpochV2NBX,
			EpochRunner:                 epochData.EpochRunner,
			EpochChallenger:             epochData.EpochChallenger,
			V2VRXTotalChallenges:        epochData.V2VRXTotalChallenges,
			V2VBXTotalChallenges:        epochData.V2VBXTotalChallenges,
			V2NBXTotalChallenges:        epochData.V2NBXTotalChallenges,
			RunnerTotalChallenges:       epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges:   epochData.ChallengerTotalChallenges,
			ChallengerPerChallengeValue: epochData.ChallengerPerChallengeValue,
			V2NBXPerChallengeValue:      epochData.V2NBXPerChallengeValue,
			RunnerPerChallengeValue:     epochData.RunnerPerChallengeValue,
			InitialPerChallengeValue:    epochData.InitialPerChallengeValue,
			TotalChallengesPrevDay:      epochData.TotalChallengesPrevDay,
		}
		k.SetEpochData(ctx, newEpochData)

	case constants.V2NBX:
		if logger != nil {
			logger.Info("updating v2n-bx epochValue.", "UpdateEpochRewards")
		}
		epochV2NBXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2NBX)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[UpdateEpochRewards][ParseCoinsNormalized] failed. Amount: [ %T ] couldn't be parsed. Error: [ %T ]", epochV2NBXCoin, err)
		}

		newEpochV2NBXCoin := epochV2NBXCoin.Add(rewardToSet)
		newEpochV2NBX := newEpochV2NBXCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:                 epochData.TotalEpochs,
			EpochV2VRX:                  epochData.EpochV2VRX,
			EpochV2VBX:                  epochData.EpochV2VBX,
			EpochV2NBX:                  newEpochV2NBX,
			EpochRunner:                 epochData.EpochRunner,
			EpochChallenger:             epochData.EpochChallenger,
			V2VRXTotalChallenges:        epochData.V2VRXTotalChallenges,
			V2VBXTotalChallenges:        epochData.V2VBXTotalChallenges,
			V2NBXTotalChallenges:        epochData.V2NBXTotalChallenges,
			RunnerTotalChallenges:       epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges:   epochData.ChallengerTotalChallenges,
			ChallengerPerChallengeValue: epochData.ChallengerPerChallengeValue,
			V2NBXPerChallengeValue:      epochData.V2NBXPerChallengeValue,
			RunnerPerChallengeValue:     epochData.RunnerPerChallengeValue,
			InitialPerChallengeValue:    epochData.InitialPerChallengeValue,
			TotalChallengesPrevDay:      epochData.TotalChallengesPrevDay,
		}
		k.SetEpochData(ctx, newEpochData)

	case constants.Runner:

		if logger != nil {
			logger.Info("updating Runner in epoch data.", "UpdateEpochRewards")
		}

		epochRunnerCoin, err := sdk.ParseCoinNormalized(epochData.EpochRunner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[UpdateEpochRewards][ParseCoinsNormalized] failed. Amount: [ %T ] couldn't be parsed. Error: [ %T ]", epochRunnerCoin, err)

		}

		newEpochRunnerCoin := epochRunnerCoin.Add(rewardToSet)
		newEpochRunner := newEpochRunnerCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:                 epochData.TotalEpochs,
			EpochV2VRX:                  epochData.EpochV2VRX,
			EpochV2VBX:                  epochData.EpochV2VBX,
			EpochV2NBX:                  epochData.EpochV2NBX,
			EpochRunner:                 newEpochRunner,
			EpochChallenger:             epochData.EpochChallenger,
			V2VRXTotalChallenges:        epochData.V2VRXTotalChallenges,
			V2VBXTotalChallenges:        epochData.V2VBXTotalChallenges,
			V2NBXTotalChallenges:        epochData.V2NBXTotalChallenges,
			RunnerTotalChallenges:       epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges:   epochData.ChallengerTotalChallenges,
			ChallengerPerChallengeValue: epochData.ChallengerPerChallengeValue,
			V2NBXPerChallengeValue:      epochData.V2NBXPerChallengeValue,
			RunnerPerChallengeValue:     epochData.RunnerPerChallengeValue,
			InitialPerChallengeValue:    epochData.InitialPerChallengeValue,
			TotalChallengesPrevDay:      epochData.TotalChallengesPrevDay,
		}

		k.SetEpochData(ctx, newEpochData)

	case constants.Challenger:

		if logger != nil {
			logger.Info("updating challenger in epoch data.", "UpdateEpochRewards")
		}

		epochChallengerCoin, err := sdk.ParseCoinNormalized(epochData.EpochChallenger)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[UpdateEpochRewards][ParseCoinsNormalized] failed. Amount: [ %T ] couldn't be parsed. Error: [ %T ]", epochChallengerCoin, err)

		}
		newEpochChallengerCoin := epochChallengerCoin.Add(rewardToSet)
		newEpochChallenger := newEpochChallengerCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:                 epochData.TotalEpochs,
			EpochV2VRX:                  epochData.EpochV2VRX,
			EpochV2VBX:                  epochData.EpochV2VBX,
			EpochV2NBX:                  epochData.EpochV2NBX,
			EpochRunner:                 epochData.EpochRunner,
			EpochChallenger:             newEpochChallenger,
			V2VRXTotalChallenges:        epochData.V2VRXTotalChallenges,
			V2VBXTotalChallenges:        epochData.V2VBXTotalChallenges,
			V2NBXTotalChallenges:        epochData.V2NBXTotalChallenges,
			RunnerTotalChallenges:       epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges:   epochData.ChallengerTotalChallenges,
			ChallengerPerChallengeValue: epochData.ChallengerPerChallengeValue,
			V2NBXPerChallengeValue:      epochData.V2NBXPerChallengeValue,
			RunnerPerChallengeValue:     epochData.RunnerPerChallengeValue,
			InitialPerChallengeValue:    epochData.InitialPerChallengeValue,
			TotalChallengesPrevDay:      epochData.TotalChallengesPrevDay,
		}
		k.SetEpochData(ctx, newEpochData)

	case constants.V2NChallenge:
		if logger != nil {
			logger.Info("updating V2NChallenge in epoch data.", "UpdateEpochRewards")
		}
		epochCnt := epochData.ChallengerTotalChallenges + 1

		totalChallengeCount := epochData.TotalChallengesPrevDay + 1

		challengeCountBlock := epochData.ChallengerLastBlockChallenges + 1

		newEpochData := types.EpochData{
			TotalEpochs:                   epochData.TotalEpochs,
			EpochV2VRX:                    epochData.EpochV2VRX,
			EpochV2VBX:                    epochData.EpochV2VBX,
			EpochV2NBX:                    epochData.EpochV2NBX,
			EpochRunner:                   epochData.EpochRunner,
			EpochChallenger:               epochData.EpochChallenger,
			V2VRXTotalChallenges:          epochData.V2VRXTotalChallenges,
			V2VBXTotalChallenges:          epochData.V2VBXTotalChallenges,
			V2NBXTotalChallenges:          epochCnt,
			RunnerTotalChallenges:         epochCnt,
			ChallengerTotalChallenges:     epochCnt,
			V2VRXLastBlockChallenges:      epochData.V2VRXLastBlockChallenges,
			V2VBXLastBlockChallenges:      epochData.V2VBXLastBlockChallenges,
			V2NBXLastBlockChallenges:      challengeCountBlock,
			RunnerLastBlockChallenges:     challengeCountBlock,
			ChallengerLastBlockChallenges: challengeCountBlock,
			ChallengerPerChallengeValue:   epochData.ChallengerPerChallengeValue,
			V2NBXPerChallengeValue:        epochData.V2NBXPerChallengeValue,
			RunnerPerChallengeValue:       epochData.RunnerPerChallengeValue,
			InitialPerChallengeValue:      epochData.InitialPerChallengeValue,
			TotalChallengesPrevDay:        totalChallengeCount,
			V2VBXPerChallengeValue:        epochData.V2VBXPerChallengeValue,
			V2VRXPerChallengeValue:        epochData.V2VRXPerChallengeValue,
		}
		k.SetEpochData(ctx, newEpochData)

	default:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[UpdateEpochRewards] failed. Client type is not valid.")
	}

	log.Println("############## End of Update Epoch Rewards ##############")

	return nil
}

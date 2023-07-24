package keeper

import (
	"soarchain/x/poa/constants"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) UpdateEpochRewards(ctx sdk.Context, clientType string, rewardToSet sdk.Coin) error {

	epochData, isFound := k.GetEpochData(ctx)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateEpochRewards][GetEpochData] failed. Epoch data is not found!")
	}

	switch clientType {

	case constants.V2VRX:
		// Parse the current value into a sdk.Coin
		epochV2VRXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2VRX)
		if err != nil {
			return err
		}

		// Addition rewardToSet
		newEpochV2VRXCoin := epochV2VRXCoin.Add(rewardToSet)

		// Convert the result back to a string representation
		newEpochV2VRX := newEpochV2VRXCoin.String()

		// Create a new EpochData object with the updated value
		newEpochData := types.EpochData{
			TotalEpochs:               epochData.TotalEpochs,
			EpochV2VRX:                newEpochV2VRX,
			EpochV2VBX:                epochData.EpochV2VBX,
			EpochV2NBX:                epochData.EpochV2NBX,
			EpochRunner:               epochData.EpochRunner,
			EpochChallenger:           epochData.EpochChallenger,
			V2VRXtotalChallenges:      epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:      epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:      epochData.V2NBXtotalChallenges,
			RunnerTotalChallenges:     epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges: epochData.ChallengerTotalChallenges,
		}

		// Store the updated epoch data
		k.SetEpochData(ctx, newEpochData)

	case constants.V2VBX:
		epochV2VBXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2VBX)
		if err != nil {
			return err
		}
		newEpochV2VBXCoin := epochV2VBXCoin.Add(rewardToSet)
		newEpochV2VBX := newEpochV2VBXCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:               epochData.TotalEpochs,
			EpochV2VRX:                epochData.EpochV2VRX,
			EpochV2VBX:                newEpochV2VBX,
			EpochV2NBX:                epochData.EpochV2NBX,
			EpochRunner:               epochData.EpochRunner,
			EpochChallenger:           epochData.EpochChallenger,
			V2VRXtotalChallenges:      epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:      epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:      epochData.V2NBXtotalChallenges,
			RunnerTotalChallenges:     epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges: epochData.ChallengerTotalChallenges,
		}
		k.SetEpochData(ctx, newEpochData)

	case constants.V2NBX:
		epochV2NBXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2NBX)
		if err != nil {
			return err
		}
		newEpochV2NBXCoin := epochV2NBXCoin.Add(rewardToSet)
		newEpochV2NBX := newEpochV2NBXCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:               epochData.TotalEpochs,
			EpochV2VRX:                epochData.EpochV2VRX,
			EpochV2VBX:                epochData.EpochV2VBX,
			EpochV2NBX:                newEpochV2NBX,
			EpochRunner:               epochData.EpochRunner,
			EpochChallenger:           epochData.EpochChallenger,
			V2VRXtotalChallenges:      epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:      epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:      epochData.V2NBXtotalChallenges,
			RunnerTotalChallenges:     epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges: epochData.ChallengerTotalChallenges,
		}
		k.SetEpochData(ctx, newEpochData)

	case constants.Runner:
		epochRunnerCoin, err := sdk.ParseCoinNormalized(epochData.EpochRunner)
		if err != nil {
			return err
		}
		newEpochRunnerCoin := epochRunnerCoin.Add(rewardToSet)
		newEpochRunner := newEpochRunnerCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:               epochData.TotalEpochs,
			EpochV2VRX:                epochData.EpochV2VRX,
			EpochV2VBX:                epochData.EpochV2VBX,
			EpochV2NBX:                epochData.EpochV2NBX,
			EpochRunner:               newEpochRunner,
			EpochChallenger:           epochData.EpochChallenger,
			V2VRXtotalChallenges:      epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:      epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:      epochData.V2NBXtotalChallenges,
			RunnerTotalChallenges:     epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges: epochData.ChallengerTotalChallenges,
		}

		k.SetEpochData(ctx, newEpochData)

	case constants.Challenger:
		epochChallengerCoin, err := sdk.ParseCoinNormalized(epochData.EpochChallenger)
		if err != nil {
			return err
		}
		newEpochChallengerCoin := epochChallengerCoin.Add(rewardToSet)
		newEpochChallenger := newEpochChallengerCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs:               epochData.TotalEpochs,
			EpochV2VRX:                epochData.EpochV2VRX,
			EpochV2VBX:                epochData.EpochV2VBX,
			EpochV2NBX:                epochData.EpochV2NBX,
			EpochRunner:               epochData.EpochRunner,
			EpochChallenger:           newEpochChallenger,
			V2VRXtotalChallenges:      epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:      epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:      epochData.V2NBXtotalChallenges,
			RunnerTotalChallenges:     epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges: epochData.ChallengerTotalChallenges,
		}
		k.SetEpochData(ctx, newEpochData)

	case constants.V2NChallenge:

		epochCnt := epochData.ChallengerTotalChallenges
		newEpochCnt := epochCnt + 1

		newEpochData := types.EpochData{
			TotalEpochs:                   epochData.TotalEpochs,
			EpochV2VRX:                    epochData.EpochV2VRX,
			EpochV2VBX:                    epochData.EpochV2VBX,
			EpochV2NBX:                    epochData.EpochV2NBX,
			EpochRunner:                   epochData.EpochRunner,
			EpochChallenger:               epochData.EpochChallenger,
			V2VRXtotalChallenges:          epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:          epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:          newEpochCnt,
			RunnerTotalChallenges:         newEpochCnt,
			ChallengerTotalChallenges:     newEpochCnt,
			V2VRXLastBlockChallenges:      epochData.V2VRXLastBlockChallenges,
			V2VBXLastBlockChallenges:      epochData.V2VBXLastBlockChallenges,
			V2NBXLastBlockChallenges:      newEpochCnt,
			RunnerLastBlockChallenges:     newEpochCnt,
			ChallengerLastBlockChallenges: newEpochCnt,
		}
		k.SetEpochData(ctx, newEpochData)

	default:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[UpdateEpochRewards] failed. Client type is not valid.")
	}

	return nil
}

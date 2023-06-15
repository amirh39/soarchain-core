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


	case "v2v-rx":
		// Parse the current value into a sdk.Coin
		epochV2VRXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2VRX)
		if err != nil {
			return err
		}

		// Subtract rewardToSet from epochV2VRXCoin
		newEpochV2VRXCoin := epochV2VRXCoin.Sub(rewardToSet)

		// Convert the result back to a string representation
		newEpochV2VRX := newEpochV2VRXCoin.String()

		// Create a new EpochData object with the updated value

		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  newEpochV2VRX,
			EpochV2VBX:  epochData.EpochV2VBX,
			EpochV2NBX:  epochData.EpochV2NBX,
			EpochRunner: epochData.EpochRunner,
		}

		// Store the updated epoch data
		k.SetEpochData(ctx, newEpochData)


	case "v2v-bx":
		epochV2VBXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2VBX)
		if err != nil {
			return err
		}
		newEpochV2VBXCoin := epochV2VBXCoin.Sub(rewardToSet)
		newEpochV2VBX := newEpochV2VBXCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  epochData.EpochV2VRX,
			EpochV2VBX:  newEpochV2VBX,
			EpochV2NBX:  epochData.EpochV2NBX,
			EpochRunner: epochData.EpochRunner,
		}
		k.SetEpochData(ctx, newEpochData)




	case "v2n-bx":
		epochV2NBXCoin, err := sdk.ParseCoinNormalized(epochData.EpochV2NBX)
		if err != nil {
			return err
		}
		newEpochV2NBXCoin := epochV2NBXCoin.Sub(rewardToSet)
		newEpochV2NBX := newEpochV2NBXCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  epochData.EpochV2VRX,
			EpochV2VBX:  epochData.EpochV2VBX,
			EpochV2NBX:  newEpochV2NBX,
			EpochRunner: epochData.EpochRunner,
		}
		k.SetEpochData(ctx, newEpochData)

	case "runner":
		epochRunnerCoin, err := sdk.ParseCoinNormalized(epochData.EpochRunner)
		if err != nil {
			return err
		}
		newEpochRunnerCoin := epochRunnerCoin.Sub(rewardToSet)
		newEpochRunner := newEpochRunnerCoin.String()

		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  epochData.EpochV2VRX,
			EpochV2VBX:  epochData.EpochV2VBX,
			EpochV2NBX:  epochData.EpochV2NBX,
			EpochRunner: newEpochRunner,
		}
		k.SetEpochData(ctx, newEpochData)

	default:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[UpdateEpochRewards] failed. Client type is not valid.")
	}

	return nil
}

package keeper

import (
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
		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  rewardToSet.String(),
			EpochV2VBX:  epochData.EpochV2VBX,
			EpochV2NBX:  epochData.EpochV2NBX,
			EpochRunner: epochData.EpochRunner,
		}
		k.SetEpochData(ctx, newEpochData)

	case "v2v-bx":
		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  epochData.EpochV2VRX,
			EpochV2VBX:  rewardToSet.String(),
			EpochV2NBX:  epochData.EpochV2NBX,
			EpochRunner: epochData.EpochRunner,
		}
		k.SetEpochData(ctx, newEpochData)

	case "v2n-bx":
		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  epochData.EpochV2VRX,
			EpochV2VBX:  epochData.EpochV2VBX,
			EpochV2NBX:  rewardToSet.String(),
			EpochRunner: epochData.EpochRunner,
		}
		k.SetEpochData(ctx, newEpochData)

	case "runner":
		newEpochData := types.EpochData{
			TotalEpochs: epochData.TotalEpochs,
			EpochV2VRX:  epochData.EpochV2VRX,
			EpochV2VBX:  epochData.EpochV2VBX,
			EpochV2NBX:  epochData.EpochV2NBX,
			EpochRunner: rewardToSet.String(),
		}
		k.SetEpochData(ctx, newEpochData)

	default:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[UpdateEpochRewards] failed. Client type is not valid.")
	}

	return nil
}

package keeper

import (
	"context"
	"log"
	"soarchain/x/dpr/types"
	"soarchain/x/dpr/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ActivateDpr(goCtx context.Context, msg *types.MsgActivateDpr) (*types.MsgActivateDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Activation of DPR Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ActivateDpr][GetDpr] failed. DPR not registered.")
	}
	if dpr.Creator != msg.Sender {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[ActivateDpr] failed. There is no valid owner for this DPR [ %T ]", msg.DprId)
	}

	if logger != nil {
		logger.Info("A Valid DPR is found successfully", "transaction", "ActivateDpr")
	}

	if dpr.Status == 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[ActivateDpr] failed. DPR is already active.")
	}

	dprEndTime, err := utility.CalculateDPREndTime(ctx.BlockHeader().Time, int(dpr.Duration))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[ActivateDpr][CalculateDPREndTime] failed. End time of the DPR couldn't calculated.")
	}
	// Convert dprEndTime from int64 to string
	dprEndTimeStr := strconv.FormatInt(dprEndTime, 10)

	epochData, isFound := k.epochKeeper.GetEpochData(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ActivateDpr][GetEpochData] failed. Epoch data is not found!")
	}
	//Save dpr into storage
	newDpr := types.Dpr{
		Id:             dpr.Id,
		Creator:        dpr.Creator,
		SupportedPIDs:  dpr.SupportedPIDs,
		Status:         1,
		Duration:       dpr.Duration,
		DprEndTime:     dprEndTimeStr,
		DprStartEpoch:  epochData.TotalEpochs,
		DprBudget:      dpr.DprBudget,
		MaxClientCount: dpr.MaxClientCount,
		ClientCounter:  dpr.ClientCounter,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Dpr activation successfully Done.", "transaction", "ActivateDpr")
	}

	log.Println("############## End of activation a dpr Transaction ##############")

	return &types.MsgActivateDprResponse{}, nil
}

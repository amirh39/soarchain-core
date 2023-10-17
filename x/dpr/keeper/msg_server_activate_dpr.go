package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"
	"soarchain/x/dpr/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ActivateDpr(goCtx context.Context, msg *types.MsgActivateDpr) (*types.MsgActivateDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Activation a dpr Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ActivateDpr][GetDpr] failed. Dpr not registered.")
	}
	if dpr.Creator != msg.Sender {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[ActivateDpr] failed. There is no valid owner for this DPR [ %T ]", msg.DprId)
	}

	if logger != nil {
		logger.Info("A Valid DPR is found successfully", "transaction", "ActivateDpr")
	}

	if dpr.IsActive {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[ActivateDpr] failed. DPR is already active.")
	}

	if len(dpr.ClientPubkeys) == 0 || dpr.ClientPubkeys == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[ActivateDpr] failed. There is no client to activate Dpr.")
	}

	dprEndTime, err := utility.CalculateDPREndTime(ctx.BlockHeader().Time, int(dpr.Duration))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[CalculateDPREndTime] failed. End time of the DPR couldn't calculated.")
	}

	//Save dpr into storage
	newDpr := types.Dpr{
		Id:            dpr.Id,
		Creator:       dpr.Creator,
		SupportedPIDs: dpr.SupportedPIDs,
		IsActive:      true,
		ClientPubkeys: dpr.ClientPubkeys,
		Duration:      dpr.Duration,
		DPRendTime:    dprEndTime,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Dpr activation successfully Done.", "transaction", "ActivateDpr")
	}

	log.Println("############## End of activation a dpr Transaction ##############")

	return &types.MsgActivateDprResponse{}, nil
}
package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeactivateDpr(goCtx context.Context, msg *types.MsgDeactivateDpr) (*types.MsgDeactivateDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Deactivation a dpr Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDpr][GetDpr] failed. Dpr not registered.")
	}
	if dpr.Creator != msg.Sender {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[DeactivateDpr] failed. There is no valid owner for this DPR [ %T ]", msg.DprId)
	}

	if logger != nil {
		logger.Info("A Valid DPR is found successfully", "transaction", "DeactivateDpr")
	}

	if !dpr.IsActive {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[DeactivateDpr] failed. DPR is already deactivate.")
	}

	if len(dpr.ClientPubkeys) == 0 || dpr.ClientPubkeys == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[DeactivateDpr] failed. There is no client to activate Dpr.")
	}

	var newDuration uint64
	if msg.Duration != 0 {
		newDuration = msg.Duration
	} else {
		newDuration = dpr.Duration
	}

	//Save dpr into storage
	newDpr := types.Dpr{
		Id:                            dpr.Id,
		Creator:                       dpr.Creator,
		PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
		IsActive:                      false,
		Vin:                           dpr.Vin,
		ClientPubkeys:                 dpr.ClientPubkeys,
		Duration:                      newDuration,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Dpr Deactivation successfully Done.", "transaction", "DeactivateDpr")
	}

	log.Println("############## End of Deactivation a dpr Transaction ##############")

	return &types.MsgDeactivateDprResponse{}, nil
}

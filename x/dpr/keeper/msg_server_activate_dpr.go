package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ActivateDpr(goCtx context.Context, msg *types.MsgActivateDpr) (*types.MsgActivateDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Activation a dpr Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ActivateDpr][GetDpr] failed. There is no eligible client to serve this DPR.")
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

	//Save dpr into storage
	newDpr := types.Dpr{
		Id:                            dpr.Id,
		Creator:                       dpr.Creator,
		PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
		IsActive:                      true,
		Vin:                           dpr.Vin,
		ClientPubkeys:                 dpr.ClientPubkeys,
		LengthOfDpr:                   dpr.LengthOfDpr,
	}

	k.SetDpr(ctx, newDpr)

	// xx := k.GetAllDpr(ctx)

	if logger != nil {
		logger.Info("Dpr activation successfully Done.", "transaction", "ActivateDpr")
	}

	log.Println("############## End of activation a dpr Transaction ##############")

	return &types.MsgActivateDprResponse{}, nil
}

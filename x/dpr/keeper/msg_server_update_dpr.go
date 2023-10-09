package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateDpr(goCtx context.Context, msg *types.MsgUpdateDpr) (*types.MsgUpdateDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDpr][GetDpr] failed. There is no valid Dpr.")
	}

	if msg.Sender != dpr.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDpr] failed. Dpr not belong to the sender.")
	}

	//Update dpr into storage
	newDpr := types.Dpr{
		Id:            msg.DprId,
		Creator:       dpr.Creator,
		SupportedPIDs: msg.SupportedPIDs,
		IsActive:      dpr.IsActive,
		Vin:           dpr.Vin,
		ClientPubkeys: dpr.ClientPubkeys,
		Duration:      msg.Duration,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Updating Dpr is successfully Done.", "transaction", "UpdateDpr")
	}

	log.Println("############## End of Updating dpr Transaction ##############")

	return &types.MsgUpdateDprResponse{}, nil
}

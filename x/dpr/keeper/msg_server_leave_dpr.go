package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LeaveDpr(goCtx context.Context, msg *types.MsgLeaveDpr) (*types.MsgLeaveDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a dpr Transaction Started ##############")

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "LeaveDpr")
	}

	_, found := k.didKeeper.GetDidDocumentByPubkey(ctx, msg.PubKey)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetDidDocumentByPubkey] failed. The sender is not owner of the device.")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)

	var eligibility bool = false
	var newpubkeysList = []string{}

	for _, pubkey := range dpr.ClientPubkeys {
		if pubkey != msg.PubKey {
			newpubkeysList = append(newpubkeysList, pubkey)
		} else {
			eligibility = true
		}
	}

	if !eligibility {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr] failed. The sender is not owner of the DPR.")
	}

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:                            dpr.Id,
		Creator:                       dpr.Creator,
		PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
		IsActive:                      dpr.IsActive,
		Vin:                           dpr.Vin,
		ClientPubkeys:                 newpubkeysList,
		LengthOfDpr:                   dpr.LengthOfDpr,
	}

	k.SetDpr(ctx, newDpr)

	// xx := k.GetAllDpr(ctx)

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "LeaveDpr", "dpr-objects")
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgLeaveDprResponse{}, nil
}

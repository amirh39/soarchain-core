package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnterDpr(goCtx context.Context, msg *types.MsgEnterDpr) (*types.MsgEnterDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Entering a dpr Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. There is no eligible client to serve this DPR.")
	}

	did, eligible := k.didKeeper.GetEligibleDidByPubkey(ctx, msg.PubKey)
	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetEligibleDidByPubkey] failed. There is no eligible client to serve this DPR.")
	}

	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "EnterDpr")
	}

	clientPubKey := []string{}
	clientPubKey = append(clientPubKey, msg.PubKey)

	vin := []string{}
	vin = append(vin, did.Document.Vin)

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:                            dpr.Id,
		Creator:                       dpr.Creator,
		PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
		IsActive:                      dpr.IsActive,
		Vin:                           vin,
		ClientPubkeys:                 clientPubKey,
		Duration:                      dpr.Duration,
	}

	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Dpr is entered to the Dpr successfully", "transaction", "EnterDpr")
	}

	log.Println("############## End of Enter dpr Transaction ##############")

	return &types.MsgEnterDprResponse{}, nil
}

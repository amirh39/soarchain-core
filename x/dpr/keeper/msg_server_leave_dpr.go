package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func remainedPubKeys(pubkey string, clientPubkeys []string) []string {
	var newPubkeys = []string{}
	for _, clientPubkey := range clientPubkeys {
		if clientPubkey != pubkey {
			newPubkeys = append(newPubkeys, clientPubkey)
		}
	}
	return newPubkeys
}

func remainedVins(vin string, vins []string) []string {
	var newVins = []string{}
	for _, dprVin := range vins {
		if dprVin != vin {
			newVins = append(newVins, dprVin)
		}
	}
	return newVins
}

func (k msgServer) LeaveDpr(goCtx context.Context, msg *types.MsgLeaveDpr) (*types.MsgLeaveDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Leaving a dpr Transaction is Started ##############")

	did, eligible := k.didKeeper.GetEligibleDidByPubkey(ctx, msg.PubKey)
	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetEligibleDidByPubkey] failed. The sender is not eligible for the DPR.")
	}

	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "EnterDpr")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. Dpr not registered.")
	}

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:                            dpr.Id,
		Creator:                       dpr.Creator,
		PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
		IsActive:                      dpr.IsActive,
		Vin:                           remainedVins(did.Document.Vin, dpr.Vin),
		ClientPubkeys:                 remainedPubKeys(msg.PubKey, dpr.ClientPubkeys),
		Duration:                      dpr.Duration,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "LeaveDpr", "dpr-objects")
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgLeaveDprResponse{}, nil
}

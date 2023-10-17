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

func (k msgServer) LeaveDpr(goCtx context.Context, msg *types.MsgLeaveDpr) (*types.MsgLeaveDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Leaving a dpr Transaction is Started ##############")

	reputation, found := k.poaKeeper.GetReputationsByAddress(ctx, msg.Sender)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetReputationsByAddress] failed. Only motus owner can send the leaveDPR transaction.")
	}

	_, eligible := k.didKeeper.GetClientDid(ctx, msg.Sender)
	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetClientDid] failed. Client not in this DPR.")
	}
	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "LeaveDpr")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. Dpr not registered.")
	}

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:            dpr.Id,
		Creator:       dpr.Creator,
		SupportedPIDs: dpr.SupportedPIDs,
		IsActive:      dpr.IsActive,
		ClientPubkeys: remainedPubKeys(reputation.PubKey, dpr.ClientPubkeys),
		Duration:      dpr.Duration,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "LeaveDpr", "dpr-objects")
	}

	log.Println("############## End of Leaving dpr Transaction ##############")

	return &types.MsgLeaveDprResponse{}, nil
}
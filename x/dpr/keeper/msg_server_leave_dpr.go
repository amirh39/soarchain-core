package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func remainedPubKeys(pubkey string, clientPubkeys []string) []string {
	// Convert the slice into a map
	pubkeyMap := make(map[string]struct{})
	for _, clientPubkey := range clientPubkeys {
		pubkeyMap[clientPubkey] = struct{}{}
	}

	// Remove the pubkey from the map
	delete(pubkeyMap, pubkey)

	// Convert the map back into a slice struct
	var newPubkeys []string
	for clientPubkey := range pubkeyMap {
		newPubkeys = append(newPubkeys, clientPubkey)
	}

	return newPubkeys
}

func (k msgServer) LeaveDpr(goCtx context.Context, msg *types.MsgLeaveDpr) (*types.MsgLeaveDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Leaving from DPR Transaction is Started ##############")

	reputation, found := k.poaKeeper.GetReputationsByAddress(ctx, msg.Sender)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetReputationsByAddress] failed. Only motus owner can send the leaveDPR transaction.")
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetDpr] failed. DPR not registered.")
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
		logger.Info("DPR is valid", "transaction", "LeaveDpr", "dpr-objects")
	}

	log.Println("############## End of Leaving DPR Transaction ##############")

	return &types.MsgLeaveDprResponse{}, nil
}

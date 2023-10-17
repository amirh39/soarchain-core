package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"
	utility "soarchain/x/dpr/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnterDpr(goCtx context.Context, msg *types.MsgEnterDpr) (*types.MsgEnterDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Entering a dpr Transaction is Started ##############")

	reputation, found := k.poaKeeper.GetReputationsByAddress(ctx, msg.Sender)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetReputationsByAddress] failed. Only motus owner can send the joinDPR transaction.")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. There is no DPR with this DPRid.")
	}

	//TODO: create a function in utils
	for _, pubKey := range dpr.ClientPubkeys {
		if reputation.PubKey == pubKey {
			// Return an error if a match is found
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][REFACTOR] failed. PubKey already exists. Your device has already registered to this DPR.")
		}
	}

	did, eligible := k.didKeeper.GetClientDid(ctx, msg.Sender)
	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetClientDid] failed. There is no eligible client to serve this DPR.")
	}

	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "EnterDpr")
	}

	eligible, err := utility.ArePIDsSupported(did.SupportedPIDs, dpr.SupportedPIDs)
	if !eligible {
		return nil, sdkerrors.Wrap(err, "[EnterDpr][ArePIDsSupported] failed. Client's PID's are not supporting the DPR.")
	}

	if logger != nil {
		logger.Info("Client's PID's are supporting the DPR", "transaction", "EnterDpr")
	}

	// Initialize a slice to store client public keys
	var clientPubKeys []string
	clientPubKeys = dpr.ClientPubkeys
	// Function to add a new public key
	clientPubKeys = append(clientPubKeys, reputation.PubKey)

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:            dpr.Id,
		Creator:       dpr.Creator,
		SupportedPIDs: dpr.SupportedPIDs,
		IsActive:      dpr.IsActive,
		ClientPubkeys: clientPubKeys,
		Duration:      dpr.Duration,
	}

	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Client is entered to the Dpr successfully", "transaction", "EnterDpr")
	}

	log.Println("############## End of Enter dpr Transaction ##############")

	return &types.MsgEnterDprResponse{}, nil
}

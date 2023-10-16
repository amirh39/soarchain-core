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

	_, found := k.poaKeeper.GetReputationsByAddress(ctx, msg.Sender)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetEligibleDidByPubkey] failed. Only motus owner can send the leaveDPR transaction.")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. There is no DPR with this DPRid.")
	}

	did, eligible := k.didKeeper.GetClientDid(ctx, msg.Sender)
	log.Println("BEFORE tryyyYYYYYYYYYYYYYYY")
	if !eligible {

		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetClientDid] failed. There is no eligible client to serve this DPR.")
	}

	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "EnterDpr")
	}

	eligible, err := utility.ArePIDsSupported(did.SupportedPIDs, dpr.SupportedPIDs)
	if eligible {
		logger.Info("Client's PID's are supporting the DPR", "transaction", "EnterDpr")
	} else {
		return nil, sdkerrors.Wrap(err, "[EnterDpr][ArePIDsSupported] failed. Client's PID's are not supporting the DPR.")
	}
	// Initialize a slice to store client public keys
	var clientPubKeys []string
	clientPubKeys = dpr.ClientPubkeys
	// Function to add a new public key
	clientPubKeys = append(clientPubKeys, msg.PubKey)

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

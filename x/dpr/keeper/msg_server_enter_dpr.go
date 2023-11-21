package keeper

import (
	"context"
	"log"

	didtypes "github.com/soar-robotics/soarchain-core/x/did/types"
	"github.com/soar-robotics/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnterDpr(goCtx context.Context, msg *types.MsgEnterDpr) (*types.MsgEnterDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Entering to a DPR Transaction is Started ##############")

	result := k.VerifyEnterDprInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[EnterDpr][VerifyDprInputs] failed. Make sure you are using valid inputs.")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. There is no DPR with this DPRid.")
	}

	if dpr.MaxClientCount <= dpr.ClientCounter {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[EnterDpr][MaxClientCount] achieved.")
	}

	did, eligible := k.didKeeper.GetClientDid(ctx, msg.Sender)
	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetClientDid] failed. There is no eligible client to serve this DPR.")
	}
	targetDprId := dpr.Id
	found = false
	for _, dprInfo := range did.DprInfos {
		if dprInfo.Id == targetDprId {
			found = true
			break
		}
	}
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[EnterDpr][GetClientDid] failed. Client is already in this DPR.")
	} else {
		// Create a new DprInfo with the dpr.Id and default claimed value
		newDprInfo := &didtypes.DprInfo{
			Id:      dpr.Id,
			Claimed: "0",
		}
		// Append the new *DprInfo to the did.DprInfos slice
		did.DprInfos = append(did.DprInfos, newDprInfo)
		k.didKeeper.SetClientDid(ctx, did)

	}

	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "EnterDpr")
	}

	eligible, err := IsCarSupportsDpr(msg.GetSupportedPIDs(), dpr.GetSupportedPIDs())
	if err != nil {
		// Handle the error appropriately, possibly by logging it or returning it up the stack
		return nil, sdkerrors.Wrap(err, "Failed to check if car supports DPR")
	}

	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Car does not support all required PIDs for the DPR")
	}

	// If we get here, it means the PIDs are eligible
	if logger != nil {
		logger.Info("Client's PIDs are supporting the DPR", "transaction", "EnterDpr")
	}

	// Save dpr into storage
	dpr.ClientCounter++
	k.SetDpr(ctx, dpr)

	if logger != nil {
		logger.Info("Client is entered to the Dpr successfully", "transaction", "EnterDpr")
	}

	log.Println("############## End of Enter dpr Transaction ##############")

	return &types.MsgEnterDprResponse{}, nil
}

package keeper

import (
	"context"
	"log"

	didtypes "soarchain/x/did/types"
	"soarchain/x/dpr/types"
	utility "soarchain/x/dpr/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnterDpr(goCtx context.Context, msg *types.MsgEnterDpr) (*types.MsgEnterDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Entering to a DPR Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetDpr] failed. There is no DPR with this DPRid.")
	}

	if dpr.MaxClientCount <= dpr.ClientCounter {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][MaxClientCount] achieved.")
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

	eligible, err := utility.ArePIDsSupported(did.SupportedPIDs, dpr.SupportedPIDs)
	if !eligible {
		return nil, sdkerrors.Wrap(err, "[EnterDpr][ArePIDsSupported] failed. Client's PID's are not supporting the DPR.")
	}

	if logger != nil {
		logger.Info("Client's PID's are supporting the DPR", "transaction", "EnterDpr")
	}

	// Save dpr into storage
	dpr.ClientCounter = dpr.ClientCounter + 1
	k.SetDpr(ctx, dpr)

	if logger != nil {
		logger.Info("Client is entered to the Dpr successfully", "transaction", "EnterDpr")
	}

	log.Println("############## End of Enter dpr Transaction ##############")

	return &types.MsgEnterDprResponse{}, nil
}

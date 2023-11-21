package keeper

import (
	"context"
	"log"

	"github.com/soar-robotics/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LeaveDpr(goCtx context.Context, msg *types.MsgLeaveDpr) (*types.MsgLeaveDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Leaving from DPR Transaction is Started ##############")

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetDpr] failed. DPR not registered.")
	}
	if dpr.Status == 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[LeaveDpr] failed. Clients cannot leave an active DPR.")
	}
	did, eligible := k.didKeeper.GetClientDid(ctx, msg.Sender)
	if !eligible {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[LeaveDpr][GetClientDid] failed. Only motus owners can leave the DPR.")
	}

	targetDprId := dpr.Id
	found = false

	// Iterate over the dprInfos slice to find the matching DprInfo entry by ID
	for i, dprInfo := range did.DprInfos {
		if dprInfo.Id == targetDprId {
			// Remove the element at index i from did.DprInfos
			did.DprInfos = append(did.DprInfos[:i], did.DprInfos[i+1:]...)
			found = true
			break
		}
	}

	if found {
		// After removing the dpr.Id, set the updated DID using the keeper
		k.didKeeper.SetClientDid(ctx, did)
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[EnterDpr][GetClientDid] failed. Client is not in this DPR.")
	}
	if logger != nil {
		logger.Info("Eligible client is found successfully", "transaction", "LeaveDpr")
	}

	// Save dpr into storage
	dpr.ClientCounter = dpr.ClientCounter - 1
	k.SetDpr(ctx, dpr)

	if logger != nil {
		logger.Info("DPR is valid", "transaction", "LeaveDpr", "dpr-objects")
	}

	log.Println("############## End of Leaving DPR Transaction ##############")

	return &types.MsgLeaveDprResponse{}, nil
}

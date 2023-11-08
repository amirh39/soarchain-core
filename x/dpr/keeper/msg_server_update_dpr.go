package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func overrideIfNonZero(current, update uint64) uint64 {
	if update != 0 {
		return update
	}
	return current
}

func overrideIfNotEmpty(current, update string) string {
	if update != "" {
		return update
	}
	return current
}

func (k msgServer) UpdateDpr(goCtx context.Context, msg *types.MsgUpdateDpr) (*types.MsgUpdateDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	dpr, found := k.GetDpr(ctx, msg.DprId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDpr][GetDpr] failed. There is no valid DPR.")
	}

	if dpr.Status == 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[UpdateDpr] failed. an active DPR cannot be modified.")
	}

	if dpr.Status == 3 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "[ActivateDpr] failed. DPR is finished.")
	}

	if msg.Sender != dpr.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDpr] failed. DPR not belong to the sender.")
	}

	// Update DPR with new or existing values
	newDpr := types.Dpr{
		Id:             dpr.Id,
		Creator:        dpr.Creator,
		SupportedPIDs:  dpr.SupportedPIDs,
		Status:         dpr.Status,
		Duration:       overrideIfNonZero(dpr.Duration, msg.Duration),
		DprEndTime:     dpr.DprEndTime,
		DprStartEpoch:  dpr.DprStartEpoch,
		DprBudget:      overrideIfNotEmpty(dpr.DprBudget, msg.DprBudget),
		MaxClientCount: overrideIfNonZero(dpr.MaxClientCount, msg.MaxClientCount),
		ClientCounter:  dpr.ClientCounter,
		Name:           dpr.Name,
	}

	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Updating Dpr is successfully Done.", "transaction", "UpdateDpr")
	}

	log.Println("############## End of Updating dpr Transaction ##############")

	return &types.MsgUpdateDprResponse{}, nil
}

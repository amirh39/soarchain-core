package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterRunner(goCtx context.Context, msg *types.MsgUnregisterRunner) (*types.MsgUnregisterRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check runner
	runner, isFoundRunner := k.GetRunner(ctx, msg.RunnerAddress)
	if !isFoundRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Runner is not registered.")
	}
	// Check runner belongs to msg.Creator's address
	if runner.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Runner doesn't belong to msg.Creator's address!")
	}

	// Query the staked amount and refund
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	stakedAmountStr := runner.StakedAmount
	stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
	transferErr2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
	if transferErr2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
	}

	// Remove runner
	k.RemoveRunner(ctx, msg.RunnerAddress)

	return &types.MsgUnregisterRunnerResponse{}, nil
}

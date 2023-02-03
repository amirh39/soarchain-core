package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterRunner(goCtx context.Context, msg *types.MsgUnregisterRunner) (*types.MsgUnregisterRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check guard
	guard, isFound := k.GetGuard(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Guard is not registered, Not authorized!")
	}
	// check runner
	runner, isFoundRunner := k.GetRunner(ctx, msg.RunnerAddress)
	if !isFoundRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Runner is not registered.")
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

	// Remove from guard
	updatedGuard := types.Guard{
		Index:         guard.Index,
		GuardId:       guard.GuardId,
		V2XChallenger: guard.V2XChallenger,
		V2NChallenger: guard.V2NChallenger,
		Runner:        &types.Runner{},
	}
	k.SetGuard(ctx, updatedGuard)

	return &types.MsgUnregisterRunnerResponse{}, nil
}

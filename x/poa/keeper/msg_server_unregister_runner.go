package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterRunner(goCtx context.Context, msg *types.MsgUnregisterRunner) (*types.MsgUnregisterRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if exists
	runner, isFound := k.GetRunner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Runner is not registered.")
	}

	// Check removal fee
	removalFee, _ := sdk.ParseCoinsNormalized("25000000soar")
	msgFee, err := sdk.ParseCoinsNormalized(msg.Fee)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
	}
	if msgFee.IsAllLT(removalFee) || !msgFee.DenomsSubsetOf(removalFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for removal.")
	}

	// Transfer fee to the protocol, then burn it
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, removalFee)
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, removalFee)

	// Query the staked amount and refund
	stakedAmountStr := runner.StakedAmount
	stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
	transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
	}

	// Remove runner
	k.RemoveRunner(ctx, msg.Creator)

	return &types.MsgUnregisterRunnerResponse{}, nil
}

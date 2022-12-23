package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterGuard(goCtx context.Context, msg *types.MsgUnregisterGuard) (*types.MsgUnregisterGuardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check guard
	guard, isFound := k.GetGuard(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Guard is not registered, Not authorized!")
	}

	// Check removal fee
	removalFee, _ := sdk.ParseCoinsNormalized("50000000soar")
	msgFee, err := sdk.ParseCoinsNormalized(msg.Fee)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
	}
	if msgFee.IsAllLT(removalFee) || !msgFee.DenomsSubsetOf(removalFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for removal.")
	}
	// Transfer fee to the protocol, then burn it
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, removalFee)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins from account to POA module!")
	}
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, removalFee)

	// check v2x challenger
	v2xChallenger, isFoundV2XChallenger := k.GetChallenger(ctx, guard.V2XChallenger.Address)
	if isFoundV2XChallenger {
		stakedAmountStr := v2xChallenger.StakedAmount
		stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
		}
		k.RemoveChallenger(ctx, guard.V2XChallenger.Address)
	}

	// check v2n challenger
	v2nChallenger, isFoundV2NChallenger := k.GetChallenger(ctx, guard.V2NChallenger.Address)
	if isFoundV2NChallenger {
		stakedAmountStr := v2nChallenger.StakedAmount
		stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
		}
		k.RemoveChallenger(ctx, guard.V2NChallenger.Address)
	}

	// check runner
	runner, isFoundRunner := k.GetRunner(ctx, guard.Runner.Address)
	if isFoundRunner {
		stakedAmountStr := runner.StakedAmount
		stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
		}
		k.RemoveRunner(ctx, guard.Runner.Address)
	}

	k.RemoveGuard(ctx, msg.Creator)

	return &types.MsgUnregisterGuardResponse{}, nil
}

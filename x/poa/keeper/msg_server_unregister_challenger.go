package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterChallenger(goCtx context.Context, msg *types.MsgUnregisterChallenger) (*types.MsgUnregisterChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check guard
	guard, isFound := k.GetGuard(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Guard is not registered, Not authorized!")
	}

	// check challenger
	challenger, isFoundChallenger := k.GetChallenger(ctx, msg.ChallengerAddress)
	if !isFoundChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Challenger is not registered.")
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
	stakedAmountStr := challenger.StakedAmount
	stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
	transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
	}

	// Remove challenger
	k.RemoveChallenger(ctx, msg.ChallengerAddress)

	// Remove from guard
	var updatedGuard types.Guard

	if challenger.Type == "v2x" {
		updatedGuard = types.Guard{
			Index:         guard.Index,
			GuardId:       guard.GuardId,
			V2XChallenger: &types.Challenger{},
			V2NChallenger: guard.V2NChallenger,
			Runner:        guard.Runner,
		}
	} else if challenger.Type == "v2n" {
		updatedGuard = types.Guard{
			Index:         guard.Index,
			GuardId:       guard.GuardId,
			V2XChallenger: guard.V2XChallenger,
			V2NChallenger: &types.Challenger{},
			Runner:        guard.Runner,
		}
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Couldn't resolve challenger type!")
	}
	k.SetGuard(ctx, updatedGuard)

	return &types.MsgUnregisterChallengerResponse{}, nil
}

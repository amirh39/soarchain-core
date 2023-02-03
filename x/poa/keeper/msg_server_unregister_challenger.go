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

	// Check challenger is belong to msg.Creator's guard
	if guard.V2NChallenger.Address != msg.ChallengerAddress || guard.V2XChallenger.Address != msg.ChallengerAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Challenger is not belong to msg.Creator's guard!")
	}

	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Query the staked amount and refund
	stakedAmountStr := challenger.StakedAmount
	stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
	transferErr2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
	if transferErr2 != nil {
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

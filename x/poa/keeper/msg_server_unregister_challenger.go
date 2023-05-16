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

	// check challenger
	challenger, isFoundChallenger := k.GetChallenger(ctx, msg.ChallengerAddress)
	if !isFoundChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Challenger is not registered.")
	}

	// Check challenger is belong to msg.Creator's address
	if challenger.Address != msg.ChallengerAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Challenger is not belong to msg.Creator's address!")
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

	return &types.MsgUnregisterChallengerResponse{}, nil
}

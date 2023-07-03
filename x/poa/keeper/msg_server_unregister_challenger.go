package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterChallenger(goCtx context.Context, msg *types.MsgUnregisterChallenger) (*types.MsgUnregisterChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Unregister Challenger Transaction Started ##############")

	// check challenger
	challenger, isFoundChallenger := k.GetChallenger(ctx, msg.ChallengerAddress)
	if !isFoundChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UnregisterChallenger][GetChallenger] failed. Challenger is not registered.")
	}

	// Check challenger is belong to msg.Creator's address
	if challenger.Address != msg.ChallengerAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Challenger is not belong to msg.Creator's address!")
	}

	msgSenderAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[UnregisterChallenger][AccAddressFromBech32] failed. Sender addres is not valid.")
	}

	if logger != nil {
		logger.Info("Checking sender address successfully done.", "transaction", "UnregisterChallenger")
	}

	// Query the staked amount and refund
	stakedAmountStr := challenger.StakedAmount
	stakedAmount, err := sdk.ParseCoinsNormalized(stakedAmountStr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterChallenger][ParseCoinsNormalized] failed. Couldn't parse the list if coins.")
	}

	transferErr2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
	if transferErr2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterChallenger][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
	}

	// Remove challenger
	k.RemoveChallenger(ctx, msg.ChallengerAddress)

	if logger != nil {
		logger.Info("Removing challenger successfully done.", "transaction", "UnregisterChallenger")
	}

	log.Println("############## End of Unregister Challenger Transaction ##############")

	return &types.MsgUnregisterChallengerResponse{}, nil
}

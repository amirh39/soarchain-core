package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterClient(goCtx context.Context, msg *types.MsgUnregisterClient) (*types.MsgUnregisterClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the client exists
	client, isFound := k.GetClient(ctx, msg.Pubkey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Client is not registered.")
	}

	// Check if authorized
	if client.Registrant != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Registrant is not recognized!")
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

	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, removalFee)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins from account to POA module!")
	}
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, removalFee)

	// Remove client
	k.RemoveClient(ctx, msg.Pubkey)

	return &types.MsgUnregisterClientResponse{}, nil
}

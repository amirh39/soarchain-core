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
	_, isFound := k.GetClient(ctx, msg.Address)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Client is not registered.")
	}

	// Check msg sender is the client, only owner of MOTUS can remove itself
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	clientAddr, _ := sdk.AccAddressFromBech32(msg.Address)
	if !(msgSenderAddress.Equals(clientAddr)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Client signature is required!")
	}

	// Check removal fee
	removalFee, _ := sdk.ParseCoinsNormalized("25soar")
	msgFee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	if removalFee.GetDenomByIndex(0) != "soar" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid coin denominator")
	}
	if msgFee.IsAllLT(removalFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for removal.")
	}

	// Transfer fee to the protocol, then burn it
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, removalFee)
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, removalFee)

	// Remove client
	k.RemoveClient(ctx, msg.Address)

	// Update Client Count
	clientCount, isFound := k.Keeper.GetTotalClients(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Client count couldn't be fetched!")
	}
	clientCount.Count--
	k.SetTotalClients(ctx, clientCount)

	return &types.MsgUnregisterClientResponse{}, nil
}

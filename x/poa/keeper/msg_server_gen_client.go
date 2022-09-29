package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Check if the client exists
	_, isFound := k.GetClient(ctx, msg.Address)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client is already registered in challenge service.")
	}

	// 2. Check registration fee
	registrationFee, _ := sdk.ParseCoinsNormalized("25soar")
	msgFee, _ := sdk.ParseCoinsNormalized(msg.Fee)

	if msgFee.IsAllLT(registrationFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for registration.")
	}

	// 3. Transfer fee to the protocol, then burn it
	clientAddr, _ := sdk.AccAddressFromBech32(msg.Address)
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, clientAddr, types.ModuleName, registrationFee)
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, registrationFee)

	// Save client into storage
	newClient := types.Client{
		Index:       clientAddr.String(),
		Address:     clientAddr.String(),
		UniqueId:    "",
		Score:       sdk.NewInt(100).String(), // Base Score
		NetEarnings: sdk.ZeroInt().String(),
	}

	k.SetClient(ctx, newClient)
	return &types.MsgGenClientResponse{}, nil
}

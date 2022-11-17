package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Check if the client exists
	_, isFound := k.GetClient(ctx, msg.Address)
	_, isFoundAsChallenger := k.GetChallenger(ctx, msg.Address)

	if isFound || isFoundAsChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client address is already registered.")
	}

	// 2. Check registration fee
	registrationFee, _ := sdk.ParseCoinsNormalized("25soar")
	msgFee, _ := sdk.ParseCoinsNormalized(msg.Fee)

	if msgFee.IsAllLT(registrationFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for registration.")
	}

	// 3. Transfer fee to the protocol, then burn it
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, registrationFee)
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, registrationFee)

	//
	clientAddr, _ := sdk.AccAddressFromBech32(msg.Address)

	// Save client into storage
	newClient := types.Client{
		Index:              clientAddr.String(),
		Address:            clientAddr.String(),
		UniqueId:           "",
		Score:              sdk.NewInt(100).String(), // Base Score
		NetEarnings:        sdk.ZeroInt().String(),
		LastTimeChallenged: sdk.ZeroInt().String(),
	}

	k.SetClient(ctx, newClient)
	return &types.MsgGenClientResponse{}, nil
}

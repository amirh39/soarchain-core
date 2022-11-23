package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/google/uuid"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Check if the client exists
	_, isFound := k.GetClient(ctx, msg.Address)
	_, isFoundAsChallenger := k.GetChallenger(ctx, msg.Address)
	_, isFoundAsRunner := k.GetRunner(ctx, msg.Address)

	if isFound || isFoundAsChallenger || isFoundAsRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client address is already registered.")
	}

	// 2. Check registration fee
	registrationFee, _ := sdk.ParseCoinsNormalized("25soar")
	msgFee, _ := sdk.ParseCoinsNormalized(msg.Fee)

	if msgFee.GetDenomByIndex(0) != "soar" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid coin denominator")
	}
	if msgFee.IsAllLT(registrationFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for registration.")
	}

	// 3. Transfer fee to the protocol, then burn it
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, registrationFee)
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, registrationFee)

	//
	clientAddr, _ := sdk.AccAddressFromBech32(msg.Address)

	// uuid
	newUUID := uuid.Must(uuid.NewRandom()).String()

	// Save client into storage
	newClient := types.Client{
		Index:              clientAddr.String(),
		Address:            clientAddr.String(),
		UniqueId:           newUUID,
		Score:              sdk.NewInt(100).String(), // Base Score
		NetEarnings:        sdk.ZeroInt().String(),
		LastTimeChallenged: sdk.ZeroInt().String(),
	}

	k.SetClient(ctx, newClient)

	// Update Client Count
	clientCount, isFound := k.Keeper.GetTotalClients(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Client count couldn't be fetched!")
	}
	clientCount.Count++
	k.SetTotalClients(ctx, clientCount)

	return &types.MsgGenClientResponse{}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"soarchain/x/poa/types"
)

func (k msgServer) CreateTotalClients(goCtx context.Context, msg *types.MsgCreateTotalClients) (*types.MsgCreateTotalClientsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTotalClients(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var totalClients = types.TotalClients{
		Creator: msg.Creator,
		Count:   msg.Count,
	}

	k.SetTotalClients(
		ctx,
		totalClients,
	)
	return &types.MsgCreateTotalClientsResponse{}, nil
}

func (k msgServer) UpdateTotalClients(goCtx context.Context, msg *types.MsgUpdateTotalClients) (*types.MsgUpdateTotalClientsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTotalClients(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var totalClients = types.TotalClients{
		Creator: msg.Creator,
		Count:   msg.Count,
	}

	k.SetTotalClients(ctx, totalClients)

	return &types.MsgUpdateTotalClientsResponse{}, nil
}

func (k msgServer) DeleteTotalClients(goCtx context.Context, msg *types.MsgDeleteTotalClients) (*types.MsgDeleteTotalClientsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTotalClients(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTotalClients(ctx)

	return &types.MsgDeleteTotalClientsResponse{}, nil
}

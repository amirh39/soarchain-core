package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetClientByAddress(goCtx context.Context, req *types.QueryGetClientByAddressRequest) (*types.QueryGetClientByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[GetClientByAddress] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	clients := k.GetAllClient(ctx)

	targetClient := types.Client{}

	for _, client := range clients {
		if req.Address == client.Address {
			targetClient = client
			break
		}
	}

	return &types.QueryGetClientByAddressResponse{Client: &targetClient}, nil
}

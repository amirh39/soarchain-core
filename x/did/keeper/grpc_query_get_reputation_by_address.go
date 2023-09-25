package keeper

import (
	"context"

	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReputationByAddress(goCtx context.Context, req *types.QueryGetReputationByAddressRequest) (*types.QueryGetReputationByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[GetClientByAddress] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	clients := k.GetAllReputation(ctx)

	targetClient := types.Reputation{}

	for _, client := range clients {
		if req.Address == client.Address {
			targetClient = client
			break
		}
	}

	return &types.QueryGetReputationByAddressResponse{Reputation: &targetClient}, nil
}

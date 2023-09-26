package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReputationByAddress(goCtx context.Context, req *types.QueryGetReputationByAddressRequest) (*types.QueryGetReputationByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[GetReputationByAddress] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reputations := k.GetAllReputation(ctx)

	targetReputation := types.Reputation{}

	for _, reputation := range reputations {
		if req.Address == reputation.Address {
			targetReputation = reputation
			break
		}
	}

	return &types.QueryGetReputationByAddressResponse{Reputation: &targetReputation}, nil
}

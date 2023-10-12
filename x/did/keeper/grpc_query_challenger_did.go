package keeper

import (
	"context"
	"log"

	"soarchain/x/did/errors"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChallengerDidAll(c context.Context, req *types.QueryAllChallengerDidRequest) (*types.QueryAllChallengerDidResponse, error) {
	log.Println("############## Fetching all challenger did is Started ##############")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(c)

	dids := k.GetAllChallengerDid(ctx)

	log.Println("############## End of fetching all challenger dids ##############")

	return &types.QueryAllChallengerDidResponse{ChallengerDid: dids}, nil
}

func (k Keeper) ChallengerDid(c context.Context, req *types.QueryGetChallengerDidRequest) (*types.QueryGetChallengerDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ChallengerDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChallengerDid(
		ctx,
		req.Address,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[ChallengerDid][GetChallengerDid] failed. Couldn't find a challenger did from the request.")
	}

	return &types.QueryGetChallengerDidResponse{ChallengerDid: val}, nil
}

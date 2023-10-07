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

	return &types.QueryAllChallengerDidResponse{ChallengerDidDocument: dids}, nil
}

func (k Keeper) ChallengerDidDocument(c context.Context, req *types.QueryGetChallengerDidRequest) (*types.QueryGetChallengerDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ChallengerDidDocument] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChallengerDidDocument(
		ctx,
		req.Id,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[ChallengerDidDocument][GetChallengerDidDocument] failed. Couldn't find a did document from the request.")
	}

	return &types.QueryGetChallengerDidResponse{ChallengerDidDocument: val}, nil
}

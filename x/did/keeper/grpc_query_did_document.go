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

func (k Keeper) DidAll(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {
	log.Println("############## Fetching all did is Started ##############")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(c)

	dids := k.GetAllDid(ctx)

	log.Println("############## End of fetching all dids ##############")

	return &types.QueryAllDidResponse{Did: dids}, nil
}

func (k Keeper) DidDocument(c context.Context, req *types.QueryGetDidRequest) (*types.QueryGetDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[DidData] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDidDocument(
		ctx,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[DidData][DidData] failed. Couldn't find a client from the request.")
	}

	return &types.QueryGetDidResponse{DidDocument: val}, nil
}

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

func (k Keeper) ClientDidAll(c context.Context, req *types.QueryAllClientDidRequest) (*types.QueryAllClientDidResponse, error) {
	log.Println("############## Fetching all client did is Started ##############")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(c)

	dids := k.GetAllClientDid(ctx)

	log.Println("############## End of fetching all client dids ##############")

	return &types.QueryAllClientDidResponse{ClientDid: dids}, nil
}

func (k Keeper) ClientDid(c context.Context, req *types.QueryGetClientDidRequest) (*types.QueryGetClientDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ClientDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClientDid(
		ctx,
		req.Id,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[ClientDid][GetClientDidDocument] failed. Couldn't find a did document from the request.")
	}

	return &types.QueryGetClientDidResponse{ClientDid: val}, nil
}

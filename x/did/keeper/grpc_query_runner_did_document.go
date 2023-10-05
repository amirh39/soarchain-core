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

func (k Keeper) RunnerDidAll(c context.Context, req *types.QueryAllRunnerDidRequest) (*types.QueryAllRunnerDidResponse, error) {
	log.Println("############## Fetching all runner did is Started ##############")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(c)

	dids := k.GetAllRunnerDid(ctx)

	log.Println("############## End of fetching all runner dids ##############")

	return &types.QueryAllRunnerDidResponse{RunnerDidDocument: dids}, nil
}

func (k Keeper) RunnerDidDocument(c context.Context, req *types.QueryGetRunnerDidRequest) (*types.QueryGetRunnerDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[RunnerDidDocument] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRunnerDidDocument(
		ctx,
		req.Id,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[RunnerDidDocument][GetRunnerDidDocument] failed. Couldn't find a did document from the request.")
	}

	return &types.QueryGetRunnerDidResponse{RunnerDidDocument: val}, nil
}

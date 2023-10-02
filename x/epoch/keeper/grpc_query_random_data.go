package keeper

import (
	"context"

	"soarchain/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RandomData(c context.Context, req *types.QueryGetRandomDataRequest) (*types.QueryGetRandomDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[RandomData] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRandomData(ctx, req.EpochNumber)
	if !found {
		return nil, status.Error(codes.NotFound, "[RandomData][GetRandomData] failed. Couldn't find data for a random data.")
	}

	return &types.QueryGetRandomDataResponse{RandomData: val}, nil
}

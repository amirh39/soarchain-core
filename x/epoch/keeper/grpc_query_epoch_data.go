package keeper

import (
	"context"

	"github.com/amirh39/soarchain-core/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EpochData(c context.Context, req *types.QueryGetEpochDataRequest) (*types.QueryGetEpochDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[EpochData] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEpochData(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "[EpochData][GetEpochData] failed. Couldn't find data for an epoch.")
	}

	return &types.QueryGetEpochDataResponse{EpochData: val}, nil
}

package keeper

import (
	"context"

	"github.com/amirh39/soarchain-core/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MasterKey(c context.Context, req *types.QueryGetMasterKeyRequest) (*types.QueryGetMasterKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMasterKey(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMasterKeyResponse{MasterKey: val}, nil
}

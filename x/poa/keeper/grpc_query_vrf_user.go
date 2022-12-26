package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"soarchain/x/poa/types"
)

func (k Keeper) VrfUserAll(c context.Context, req *types.QueryAllVrfUserRequest) (*types.QueryAllVrfUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vrfUsers []types.VrfUser
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vrfUserStore := prefix.NewStore(store, types.KeyPrefix(types.VrfUserKeyPrefix))

	pageRes, err := query.Paginate(vrfUserStore, req.Pagination, func(key []byte, value []byte) error {
		var vrfUser types.VrfUser
		if err := k.cdc.Unmarshal(value, &vrfUser); err != nil {
			return err
		}

		vrfUsers = append(vrfUsers, vrfUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVrfUserResponse{VrfUser: vrfUsers, Pagination: pageRes}, nil
}

func (k Keeper) VrfUser(c context.Context, req *types.QueryGetVrfUserRequest) (*types.QueryGetVrfUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVrfUser(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVrfUserResponse{VrfUser: val}, nil
}

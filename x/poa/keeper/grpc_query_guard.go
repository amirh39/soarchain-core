package keeper

import (
	"context"

	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GuardAll(c context.Context, req *types.QueryAllGuardRequest) (*types.QueryAllGuardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[GuardAll] failed. Invalid request.")
	}

	var guards []types.Guard
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	guardStore := prefix.NewStore(store, types.KeyPrefix(types.GuardKeyPrefix))

	pageRes, err := query.Paginate(guardStore, req.Pagination, func(key []byte, value []byte) error {
		var guard types.Guard
		if err := k.cdc.Unmarshal(value, &guard); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[GuardAll][Unmarshal] failed. Couldn't parses the guard data encoded."+err.Error())
		}

		guards = append(guards, guard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGuardResponse{Guard: guards, Pagination: pageRes}, nil
}

func (k Keeper) Guard(c context.Context, req *types.QueryGetGuardRequest) (*types.QueryGetGuardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[Guard] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGuard(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[Guard][GetGuard] failed. Couldn't find given guard byt the request.")
	}

	return &types.QueryGetGuardResponse{Guard: val}, nil
}

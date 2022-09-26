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

func (k Keeper) ClientAll(c context.Context, req *types.QueryAllClientRequest) (*types.QueryAllClientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var clients []types.Client
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.ClientKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(key []byte, value []byte) error {
		var client types.Client
		if err := k.cdc.Unmarshal(value, &client); err != nil {
			return err
		}

		clients = append(clients, client)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClientResponse{Client: clients, Pagination: pageRes}, nil
}

func (k Keeper) Client(c context.Context, req *types.QueryGetClientRequest) (*types.QueryGetClientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClient(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClientResponse{Client: val}, nil
}

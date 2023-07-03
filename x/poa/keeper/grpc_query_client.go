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

func (k Keeper) ClientAll(c context.Context, req *types.QueryAllClientRequest) (*types.QueryAllClientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ClientAll] failed. Invalid request.")
	}

	var clients []types.Client
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.ClientKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(key []byte, value []byte) error {
		var client types.Client
		if err := k.cdc.Unmarshal(value, &client); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[ClientAll][Unmarshal] failed. Couldn't parse the client data encoded."+err.Error())
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
		return nil, status.Error(codes.InvalidArgument, "[Client] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClient(
		ctx,
		req.PubKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[Client][GetClient] failed. Couldn't find a client from the request.")
	}

	return &types.QueryGetClientResponse{Client: val}, nil
}

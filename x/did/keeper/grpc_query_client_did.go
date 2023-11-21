package keeper

import (
	"context"

	"github.com/soar-robotics/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ClientDidAll(c context.Context, req *types.QueryAllClientDidRequest) (*types.QueryAllClientDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ClientDidAll] failed. Invalid request.")
	}

	var clientDids []types.ClientDid
	ctx := sdk.UnwrapSDKContext(c)

	pagination := req.Pagination
	if pagination == nil {
		pagination = &query.PageRequest{Limit: 100}
	} else if pagination.Limit == 0 || pagination.Limit > 1000 {
		pagination.Limit = 1000
	}

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.DidKeyPrefix))

	pageRes, err := query.Paginate(clientStore, pagination, func(key []byte, value []byte) error {
		var clientDid types.ClientDid
		if err := k.cdc.Unmarshal(value, &clientDid); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[ClientDidAll][Unmarshal] failed. Couldn't parse the reputation data encoded.")
		}
		clientDids = append(clientDids, clientDid)
		return nil
	})

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[ClientDidAll] failed. Couldn't find a client did.")
	}

	return &types.QueryAllClientDidResponse{ClientDid: clientDids, Pagination: pageRes}, nil
}

func (k Keeper) ClientDid(c context.Context, req *types.QueryGetClientDidRequest) (*types.QueryGetClientDidResponse, error) {
	if req == nil || req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "[ClientDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClientDid(
		ctx,
		req.Address,
	)

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "[ClientDid][GetClientDid] failed. Couldn't find a valid client did by the this address: [ %s ], Make sure address is not empty OR invalid.", req.Address)
	}

	return &types.QueryGetClientDidResponse{ClientDid: val}, nil
}

package keeper

import (
	"context"

	"soarchain/x/did/types"

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

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.DidKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(key []byte, value []byte) error {
		var clientDid types.ClientDid
		if err := k.cdc.Unmarshal(value, &clientDid); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[ClientDidAll][Unmarshal] failed. Couldn't parse the reputation data encoded.")
		}

		clientDids = append(clientDids, clientDid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClientDidResponse{ClientDid: clientDids, Pagination: pageRes}, nil
}

func (k Keeper) ClientDid(c context.Context, req *types.QueryGetClientDidRequest) (*types.QueryGetClientDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ClientDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClientDid(
		ctx,
		req.Address,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[ClientDid][GetClientDidDocument] failed. Couldn't find a did document from the request.")
	}

	return &types.QueryGetClientDidResponse{ClientDid: val}, nil
}

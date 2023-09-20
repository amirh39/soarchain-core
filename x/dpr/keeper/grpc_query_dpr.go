package keeper

import (
	"context"

	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DprAll(c context.Context, req *types.QueryAllDprRequest) (*types.QueryAllDprResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[DprAll] failed. Invalid request.")
	}

	var dprs []types.Dpr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.DprKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(_ []byte, value []byte) error {
		var dpr types.Dpr
		if err := k.cdc.Unmarshal(value, &dpr); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[DprAll][Unmarshal] failed. Couldn't parse the dpr data encoded.")
		}

		dprs = append(dprs, dpr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDprResponse{Dpr: dprs, Pagination: pageRes}, nil
}

func (k Keeper) Dpr(c context.Context, req *types.QueryGetDprRequest) (*types.QueryGetDprResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[Dpr] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDpr(
		ctx,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[Dpr][GetDpr] failed. Couldn't find a dpr from the request.")
	}

	return &types.QueryGetDprResponse{Dpr: val}, nil
}

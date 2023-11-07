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

func (k Keeper) VrfDataAll(c context.Context, req *types.QueryAllVrfDataRequest) (*types.QueryAllVrfDataResponse, error) {
	if req == nil || req.Pagination == nil {
		return nil, status.Error(codes.InvalidArgument, "[VrfDataAll] failed. Invalid request.")
	}

	var vrfDatas []types.VrfData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vrfDataStore := prefix.NewStore(store, types.KeyPrefix(types.VrfDataKeyPrefix))

	pageRes, err := query.Paginate(vrfDataStore, req.Pagination, func(key []byte, value []byte) error {
		var vrfData types.VrfData
		if err := k.cdc.Unmarshal(value, &vrfData); err != nil {
			return err
		}

		vrfDatas = append(vrfDatas, vrfData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVrfDataResponse{VrfData: vrfDatas, Pagination: pageRes}, nil
}

func (k Keeper) VrfData(c context.Context, req *types.QueryGetVrfDataRequest) (*types.QueryGetVrfDataResponse, error) {
	if req == nil || req.Index == "" {
		return nil, status.Error(codes.InvalidArgument, "[VrfData] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVrfData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[VrfData][GetVrfData] failed. Couldn't find a valid vrf for this index: [ %s ].", req.Index)
	}
	return &types.QueryGetVrfDataResponse{VrfData: val}, nil
}

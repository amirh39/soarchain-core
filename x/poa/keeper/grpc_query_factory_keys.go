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

func (k Keeper) FactoryKeysAll(c context.Context, req *types.QueryAllFactoryKeysRequest) (*types.QueryAllFactoryKeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[FactoryKeysAll] failed. Invalid request.")
	}

	var factoryKeys []types.FactoryKeys
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	factoryKeysStore := prefix.NewStore(store, types.KeyPrefix(types.FactoryKeysKey))

	pageRes, err := query.Paginate(factoryKeysStore, req.Pagination, func(key []byte, value []byte) error {
		var factoryKey types.FactoryKeys
		if err := k.cdc.Unmarshal(value, &factoryKey); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[FactoryKeysAll][Unmarshal] failed. Couldn't parses the factory data encoded.")
		}

		factoryKeys = append(factoryKeys, factoryKey)
		return nil
	})

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[FactoryKeysAll] failed. Couldn't find a valid factory key.")
	}

	return &types.QueryAllFactoryKeysResponse{FactoryKeys: factoryKeys, Pagination: pageRes}, nil
}

func (k Keeper) FactoryKeys(c context.Context, req *types.QueryGetFactoryKeysRequest) (*types.QueryGetFactoryKeysResponse, error) {
	if req == nil || req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "[FactoryKeys] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(c)
	factoryKeys, found := k.GetFactoryKeys(ctx, req.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[GetReputationByAddress][GetReputationsByAddress] failed. Couldn't find a valid reputation for this address: [ %T ].", req.Id)
	}

	return &types.QueryGetFactoryKeysResponse{FactoryKeys: factoryKeys}, nil
}

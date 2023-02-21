package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"soarchain/x/poa/types"
)

func (k Keeper) FactoryKeysAll(c context.Context, req *types.QueryAllFactoryKeysRequest) (*types.QueryAllFactoryKeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var factoryKeyss []types.FactoryKeys
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	factoryKeysStore := prefix.NewStore(store, types.KeyPrefix(types.FactoryKeysKey))

	pageRes, err := query.Paginate(factoryKeysStore, req.Pagination, func(key []byte, value []byte) error {
		var factoryKeys types.FactoryKeys
		if err := k.cdc.Unmarshal(value, &factoryKeys); err != nil {
			return err
		}

		factoryKeyss = append(factoryKeyss, factoryKeys)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFactoryKeysResponse{FactoryKeys: factoryKeyss, Pagination: pageRes}, nil
}

func (k Keeper) FactoryKeys(c context.Context, req *types.QueryGetFactoryKeysRequest) (*types.QueryGetFactoryKeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	factoryKeys, found := k.GetFactoryKeys(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFactoryKeysResponse{FactoryKeys: factoryKeys}, nil
}

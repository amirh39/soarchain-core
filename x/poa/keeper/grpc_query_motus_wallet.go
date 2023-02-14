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

func (k Keeper) MotusWalletAll(c context.Context, req *types.QueryAllMotusWalletRequest) (*types.QueryAllMotusWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var motusWallets []types.MotusWallet
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	motusWalletStore := prefix.NewStore(store, types.KeyPrefix(types.MotusWalletKeyPrefix))

	pageRes, err := query.Paginate(motusWalletStore, req.Pagination, func(key []byte, value []byte) error {
		var motusWallet types.MotusWallet
		if err := k.cdc.Unmarshal(value, &motusWallet); err != nil {
			return err
		}

		motusWallets = append(motusWallets, motusWallet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMotusWalletResponse{MotusWallet: motusWallets, Pagination: pageRes}, nil
}

func (k Keeper) MotusWallet(c context.Context, req *types.QueryGetMotusWalletRequest) (*types.QueryGetMotusWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMotusWallet(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMotusWalletResponse{MotusWallet: val}, nil
}

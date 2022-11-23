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

func (k Keeper) ChallengerByIndexAll(c context.Context, req *types.QueryAllChallengerByIndexRequest) (*types.QueryAllChallengerByIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var challengerByIndexs []types.ChallengerByIndex
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	challengerByIndexStore := prefix.NewStore(store, types.KeyPrefix(types.ChallengerByIndexKeyPrefix))

	pageRes, err := query.Paginate(challengerByIndexStore, req.Pagination, func(key []byte, value []byte) error {
		var challengerByIndex types.ChallengerByIndex
		if err := k.cdc.Unmarshal(value, &challengerByIndex); err != nil {
			return err
		}

		challengerByIndexs = append(challengerByIndexs, challengerByIndex)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChallengerByIndexResponse{ChallengerByIndex: challengerByIndexs, Pagination: pageRes}, nil
}

func (k Keeper) ChallengerByIndex(c context.Context, req *types.QueryGetChallengerByIndexRequest) (*types.QueryGetChallengerByIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChallengerByIndex(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetChallengerByIndexResponse{ChallengerByIndex: val}, nil
}

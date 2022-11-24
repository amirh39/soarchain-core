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

func (k Keeper) RunnerByIndexAll(c context.Context, req *types.QueryAllRunnerByIndexRequest) (*types.QueryAllRunnerByIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var runnerByIndexs []types.RunnerByIndex
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	runnerByIndexStore := prefix.NewStore(store, types.KeyPrefix(types.RunnerByIndexKeyPrefix))

	pageRes, err := query.Paginate(runnerByIndexStore, req.Pagination, func(key []byte, value []byte) error {
		var runnerByIndex types.RunnerByIndex
		if err := k.cdc.Unmarshal(value, &runnerByIndex); err != nil {
			return err
		}

		runnerByIndexs = append(runnerByIndexs, runnerByIndex)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRunnerByIndexResponse{RunnerByIndex: runnerByIndexs, Pagination: pageRes}, nil
}

func (k Keeper) RunnerByIndex(c context.Context, req *types.QueryGetRunnerByIndexRequest) (*types.QueryGetRunnerByIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRunnerByIndex(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRunnerByIndexResponse{RunnerByIndex: val}, nil
}

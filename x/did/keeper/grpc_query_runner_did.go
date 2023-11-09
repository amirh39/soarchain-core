package keeper

import (
	"context"

	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) RunnerDidAll(c context.Context, req *types.QueryAllRunnerDidRequest) (*types.QueryAllRunnerDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[RunnerDidAll] failed. Invalid request.")
	}

	var runnerDids []types.RunnerDid
	ctx := sdk.UnwrapSDKContext(c)

	pagination := req.Pagination
	if pagination == nil {
		pagination = &query.PageRequest{Limit: 100}
	} else if pagination.Limit == 0 || pagination.Limit > 1000 {
		pagination.Limit = 1000
	}

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.RunnerDidKeyPrefix))

	pageRes, err := query.Paginate(clientStore, pagination, func(key []byte, value []byte) error {
		var runnerDid types.RunnerDid
		if err := k.cdc.Unmarshal(value, &runnerDid); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[RunnerDidAll][Unmarshal] failed. Couldn't parse the runner did data encoded.")
		}

		runnerDids = append(runnerDids, runnerDid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRunnerDidResponse{RunnerDid: runnerDids, Pagination: pageRes}, nil
}

func (k Keeper) RunnerDid(c context.Context, req *types.QueryGetRunnerDidRequest) (*types.QueryGetRunnerDidResponse, error) {
	if req == nil || req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "[RunnerDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRunnerDid(
		ctx,
		req.Address,
	)

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[RunnerDid][GetRunnerDid] failed. Couldn't find a valid runner for this address: [ %s ] .", req.Address)
	}

	return &types.QueryGetRunnerDidResponse{RunnerDid: val}, nil
}

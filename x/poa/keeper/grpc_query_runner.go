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

func (k Keeper) RunnerAll(c context.Context, req *types.QueryAllRunnerRequest) (*types.QueryAllRunnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[RunnerAll] failed. Invalid request.")
	}

	var runners []types.Runner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	runnerStore := prefix.NewStore(store, types.KeyPrefix(types.RunnerKeyPrefix))

	pageRes, err := query.Paginate(runnerStore, req.Pagination, func(key []byte, value []byte) error {
		var runner types.Runner
		if err := k.cdc.Unmarshal(value, &runner); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[RunnerAll][Unmarshal] failed. Couldn't parses the data encoded."+err.Error())
		}

		runners = append(runners, runner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRunnerResponse{Runner: runners, Pagination: pageRes}, nil
}

func (k Keeper) Runner(c context.Context, req *types.QueryGetRunnerRequest) (*types.QueryGetRunnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[Runner] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRunner(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[Runner] failed. Couldn't find a runner.")
	}

	return &types.QueryGetRunnerResponse{Runner: val}, nil
}

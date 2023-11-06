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

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.DidKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(key []byte, value []byte) error {
		var runnerDid types.RunnerDid
		if err := k.cdc.Unmarshal(value, &runnerDid); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[RunnerDidAll][Unmarshal] failed. Couldn't parse the reputation data encoded.")
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
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[RunnerDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRunnerDid(
		ctx,
		req.Address,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[RunnerDid][GetRunnerDidDocument] failed. Couldn't find a did document from the request.")
	}

	return &types.QueryGetRunnerDidResponse{RunnerDid: val}, nil
}

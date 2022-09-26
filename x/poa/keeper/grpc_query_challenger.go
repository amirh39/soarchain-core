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

func (k Keeper) ChallengerAll(c context.Context, req *types.QueryAllChallengerRequest) (*types.QueryAllChallengerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var challengers []types.Challenger
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	challengerStore := prefix.NewStore(store, types.KeyPrefix(types.ChallengerKeyPrefix))

	pageRes, err := query.Paginate(challengerStore, req.Pagination, func(key []byte, value []byte) error {
		var challenger types.Challenger
		if err := k.cdc.Unmarshal(value, &challenger); err != nil {
			return err
		}

		challengers = append(challengers, challenger)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChallengerResponse{Challenger: challengers, Pagination: pageRes}, nil
}

func (k Keeper) Challenger(c context.Context, req *types.QueryGetChallengerRequest) (*types.QueryGetChallengerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChallenger(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetChallengerResponse{Challenger: val}, nil
}

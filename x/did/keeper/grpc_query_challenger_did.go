package keeper

import (
	"context"

	"github.com/amirh39/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) ChallengerDidAll(c context.Context, req *types.QueryAllChallengerDidRequest) (*types.QueryAllChallengerDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ChallengerDidAll] failed. Invalid request.")
	}

	var challengerDids []types.ChallengerDid
	ctx := sdk.UnwrapSDKContext(c)

	pagination := req.Pagination
	if pagination == nil {
		pagination = &query.PageRequest{Limit: 100}
	} else if pagination.Limit == 0 || pagination.Limit > 1000 {
		pagination.Limit = 1000
	}

	store := ctx.KVStore(k.storeKey)
	challengerStore := prefix.NewStore(store, types.KeyPrefix(types.ChallengerDidKeyPrefix))

	pageRes, err := query.Paginate(challengerStore, pagination, func(key []byte, value []byte) error {
		var challengerDid types.ChallengerDid
		if err := k.cdc.Unmarshal(value, &challengerDid); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[ChallengerDidAll][Unmarshal] failed. Couldn't parse the challenger did data encoded.")
		}

		challengerDids = append(challengerDids, challengerDid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChallengerDidResponse{ChallengerDid: challengerDids, Pagination: pageRes}, nil
}

func (k Keeper) ChallengerDid(c context.Context, req *types.QueryGetChallengerDidRequest) (*types.QueryGetChallengerDidResponse, error) {
	if req == nil || req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "[ChallengerDid] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChallengerDid(
		ctx,
		req.Address,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "[ChallengerDid][GetChallengerDid] failed. Couldn't find a challenger did from the request.")
	}

	return &types.QueryGetChallengerDidResponse{ChallengerDid: val}, nil
}

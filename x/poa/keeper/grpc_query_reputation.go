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

func (k Keeper) ReputationAll(c context.Context, req *types.QueryAllReputationRequest) (*types.QueryAllReputationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[ReputationAll] failed. Invalid request.")
	}

	var reputations []types.Reputation
	ctx := sdk.UnwrapSDKContext(c)

	limit := req.Pagination.GetLimit()
	if limit == 0 || limit > 100 {
		limit = 100
	}

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.ReputationKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(key []byte, value []byte) error {
		var reputation types.Reputation
		if err := k.cdc.Unmarshal(value, &reputation); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[ReputationAll][Unmarshal] failed. Couldn't parse the reputation data encoded.")
		}

		reputations = append(reputations, reputation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReputationResponse{Reputation: reputations, Pagination: pageRes}, nil
}

func (k Keeper) Reputation(c context.Context, req *types.QueryGetReputationRequest) (*types.QueryGetReputationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[Reputation] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetReputation(
		ctx,
		req.Pubkey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[Reputation][GetReputation] failed. Make sure index of the reputation is valid.")
	}

	return &types.QueryGetReputationResponse{Reputation: val}, nil
}

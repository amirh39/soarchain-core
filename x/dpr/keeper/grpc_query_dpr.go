package keeper

import (
	"context"

	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DprAll(c context.Context, req *types.QueryAllDprRequest) (*types.QueryAllDprResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[DprAll] failed. Invalid request.")
	}

	var dprs []types.Dpr
	ctx := sdk.UnwrapSDKContext(c)

	limit := req.Pagination.GetLimit()
	if limit == 0 || limit > 100 {
		limit = 100
	}

	store := ctx.KVStore(k.storeKey)
	clientStore := prefix.NewStore(store, types.KeyPrefix(types.DprKeyPrefix))

	pageRes, err := query.Paginate(clientStore, req.Pagination, func(_ []byte, value []byte) error {
		var dpr types.Dpr
		if err := k.cdc.Unmarshal(value, &dpr); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[DprAll][Unmarshal] failed. Couldn't parse the dpr data encoded.")
		}

		dprs = append(dprs, dpr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.NotFound, "[Dpr][DprAll] failed. Couldn't get dprs.")
	}

	return &types.QueryAllDprResponse{Dpr: dprs, Pagination: pageRes}, nil
}

func (k Keeper) Dpr(c context.Context, req *types.QueryGetDprRequest) (*types.QueryGetDprResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[Dpr] failed. Invalid request.")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDpr(
		ctx,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "[Dpr][GetDpr] failed. Couldn't find a dpr from the request.")
	}

	return &types.QueryGetDprResponse{Dpr: val}, nil
}

func (k Keeper) DPRsByClientPubkey(c context.Context, req *types.QueryDPRsByClientPubkeyRequest) (*types.QueryDPRsByClientPubkeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid request")
	}

	// TODO: Add a check to verify that the provided pubkey is a registered client

	var matchingDprs []*types.Dpr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	dprStore := prefix.NewStore(store, types.KeyPrefix(types.DprKeyPrefix))

	pageRes, err := query.Paginate(dprStore, req.Pagination, func(_ []byte, value []byte) error {
		var dpr types.Dpr
		if err := k.cdc.Unmarshal(value, &dpr); err != nil {
			return err
		}

		// for _, pubkey := range dpr.ClientPubkeys {
		// 	if pubkey == req.ClientPubkey {
		// 		matchingDprs = append(matchingDprs, &dpr)
		// 		break
		// 	}
		// }
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.NotFound, "[Dpr][DPRsByClientPubkey] failed. Invalid query parameters")
	}

	return &types.QueryDPRsByClientPubkeyResponse{Dpr: matchingDprs, Pagination: pageRes}, nil
}

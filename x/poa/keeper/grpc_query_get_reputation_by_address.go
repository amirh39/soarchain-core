package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReputationByAddress(goCtx context.Context, req *types.QueryGetReputationByAddressRequest) (*types.QueryGetReputationByAddressResponse, error) {
	if req == nil || req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "[GetReputationByAddress] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reputation, found := k.GetReputationsByAddress(ctx, req.Address)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[GetReputationByAddress][GetReputationsByAddress] failed. Couldn't find a valid reputation for this address: [ %s ].", req.Address)
	}

	return &types.QueryGetReputationByAddressResponse{Reputation: &reputation}, nil
}

package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetChallengerByAddress(goCtx context.Context, req *types.QueryGetChallengerByAddressRequest) (*types.QueryGetChallengerByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[GetChallengerByAddress] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, found := k.GetChallenger(ctx, req.Address)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[GetChallengerByAddress][GetChallenger] failed. Couldn't find a challenger by the address.")
	}

	return &types.QueryGetChallengerByAddressResponse{Challenger: &challenger}, nil
}

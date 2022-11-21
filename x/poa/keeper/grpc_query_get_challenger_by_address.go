package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetChallengerByAddress(goCtx context.Context, req *types.QueryGetChallengerByAddressRequest) (*types.QueryGetChallengerByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, _ := k.GetChallenger(ctx, req.Address)

	return &types.QueryGetChallengerByAddressResponse{Challenger: &challenger}, nil
}

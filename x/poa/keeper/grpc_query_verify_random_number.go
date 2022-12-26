package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyRandomNumber(goCtx context.Context, req *types.QueryVerifyRandomNumberRequest) (*types.QueryVerifyRandomNumberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	isVerified, _ := k.VerifyGeneratedNumber(ctx, req)

	return &types.QueryVerifyRandomNumberResponse{Result: isVerified}, nil
}

package keeper

import (
	"context"

	"math/rand"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRandomRunner(goCtx context.Context, req *types.QueryGetRandomRunnerRequest) (*types.QueryGetRandomRunnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	allRunner := k.GetAllRunner(ctx)

	rand.Seed(ctx.BlockTime().UnixNano())
	min := 0
	max := int(len(allRunner) - 1)
	n := rand.Intn(max-min+1) + min

	return &types.QueryGetRandomRunnerResponse{Runner: (&allRunner[n])}, nil
}

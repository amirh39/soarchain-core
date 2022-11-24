package keeper

import (
	"context"

	"math/rand"
	"strconv"

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

	// TODO: Process the query
	totalChallengers, _ := k.GetTotalRunners(ctx)

	rand.Seed(ctx.BlockTime().UnixNano())
	min := 1
	max := int(totalChallengers.Count)
	n := rand.Intn(max-min+1) + min
	indexStr := strconv.Itoa(int(n))
	randRunner, _ := k.GetRunnerByIndex(ctx, indexStr)

	return &types.QueryGetRandomRunnerResponse{Runner: (randRunner.Runner)}, nil
}

package keeper

import (
	"context"

	// "strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRandomChallenger(goCtx context.Context, req *types.QueryGetRandomChallengerRequest) (*types.QueryGetRandomChallengerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	randomNumber := k.GenerateRandomNumber(goCtx)
	challengerArr := k.GetAllChallenger(ctx)

	return &types.QueryGetRandomChallengerResponse{Challenger: (&challengerArr[randomNumber])}, nil

}

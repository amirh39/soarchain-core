package keeper

import (
	"context"

	"strconv"

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
	indexStr := strconv.Itoa(int(randomNumber))
	randChallenger, _ := k.GetChallengerByIndex(ctx, indexStr)

	return &types.QueryGetRandomChallengerResponse{Challenger: (randChallenger.Challenger)}, nil

	// for {
	// 	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(totalChallengers.Count+1)))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	n := nBig.Int64()
	// 	if n != 0 {
	// 		indexStr := strconv.Itoa(int(n))
	// 		randChallenger, _ := k.GetChallengerByIndex(ctx, indexStr)

	// 		return &types.QueryGetRandomChallengerResponse{Challenger: (randChallenger.Challenger)}, nil
	// 	}
	// }

}

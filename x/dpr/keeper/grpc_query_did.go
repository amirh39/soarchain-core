package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"
)

func (k Keeper) DprAll(c context.Context, req *types.QueryAllDprRequest) (*types.QueryAllDprResponse, error) {

	log.Println("############## Fetching all dpr is Started ##############")

	// if req == nil {
	// 	return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	// }

	//ctx := sdk.UnwrapSDKContext(c)

	//dprs := k.GetAllDpr(ctx)

	log.Println("############## End of fetching all dprs ##############")

	return &types.QueryAllDprResponse{Dpr: nil}, nil
}

func (k Keeper) Dpr(c context.Context, req *types.QueryGetDprRequest) (*types.QueryGetDprResponse, error) {

	log.Println("############## Fetching a dpr is Started ##############")

	//ctx := sdk.UnwrapSDKContext(c)

	// if req == nil {
	// 	return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	// }

	log.Println("############## End of fetching a dpr ##############")

	return &types.QueryGetDprResponse{Dpr: nil}, nil
}

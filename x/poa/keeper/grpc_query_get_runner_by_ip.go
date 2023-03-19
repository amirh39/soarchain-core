package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRunnerByIp(goCtx context.Context, req *types.QueryGetRunnerByIpRequest) (*types.QueryGetRunnerByIpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	runners := k.GetAllRunner(ctx)

	targetRunner := types.Runner{}

	for _, runner := range runners {
		if req.IpAddress == runner.IpAddr {
			targetRunner = runner
			break
		}
	}

	return &types.QueryGetRunnerByIpResponse{Runner: &targetRunner}, nil
}

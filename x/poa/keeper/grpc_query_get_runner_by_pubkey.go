package keeper

import (
	"context"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRunnerByPubKey(goCtx context.Context, req *types.QueryGetRunnerByPubKeyRequest) (*types.QueryGetRunnerByPubKeyResponse, error) {
	if req == nil || req.Pubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	runners := k.GetAllRunner(ctx)

	targetRunner := types.Runner{}

	for _, runner := range runners {
		if req.Pubkey == runner.PubKey {
			targetRunner = runner
			break
		}
	}

	if targetRunner.PubKey == "" {
		return nil, status.Error(codes.NotFound, "runner not found")
	}

	return &types.QueryGetRunnerByPubKeyResponse{Runner: &targetRunner}, nil
}

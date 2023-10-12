package keeper

import (
	"context"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRunnerDidByPubKey(goCtx context.Context, req *types.QueryGetRunnerDidByPubKeyRequest) (*types.QueryGetRunnerDidByPubKeyResponse, error) {
	if req == nil || req.Pubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	runners := k.GetAllRunnerDid(ctx)

	targetRunner := types.RunnerDid{}

	for _, runner := range runners {
		if req.Pubkey == runner.PubKey {
			targetRunner = runner
			break
		}
	}

	if targetRunner.PubKey == "" {
		return nil, status.Error(codes.NotFound, "runner not found")
	}

	return &types.QueryGetRunnerDidByPubKeyResponse{RunnerDid: &targetRunner}, nil
}

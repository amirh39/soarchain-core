package keeper

import (
	"context"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetChallengerByPubKey(goCtx context.Context, req *types.QueryGetChallengerByPubKeyRequest) (*types.QueryGetChallengerByPubKeyResponse, error) {
	if req == nil || req.PubKey == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challengers := k.GetAllChallenger(ctx)

	targetChallenger := types.Challenger{}

	for _, challenger := range challengers {
		if req.PubKey == challenger.PubKey {
			targetChallenger = challenger
			break
		}
	}

	if targetChallenger.PubKey == "" {
		return nil, status.Error(codes.NotFound, "challenger not found")
	}

	return &types.QueryGetChallengerByPubKeyResponse{Challenger: &targetChallenger}, nil
}

package keeper

import (
	"context"

	"github.com/amirh39/soarchain-core/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetChallengerDidByPubKey(goCtx context.Context, req *types.QueryGetChallengerDidByPubKeyRequest) (*types.QueryGetChallengerDidByPubKeyResponse, error) {
	if req == nil || req.Pubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "[GetChallengerDidByPubKey] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, found := k.GetChallengerDidUsingPubKey(ctx, req.Pubkey)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[GetChallengerDidByPubKey][GetRunnerDidUsingPubKey] failed. Couldn't find a valid challenger for this pubkey: [ %s ] .", req.Pubkey)
	}

	return &types.QueryGetChallengerDidByPubKeyResponse{ChallengerDid: &challenger}, nil
}

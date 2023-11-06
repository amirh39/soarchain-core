package keeper

import (
	"context"
	"soarchain/x/did/types"

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

	challengers, found := k.GetChallengerDidUsingPubKey(ctx, req.Pubkey)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[GetChallengerDidByPubKey][GetRunnerDidUsingPubKey] failed. Couldn't find a valid challenger.")
	}

	return &types.QueryGetChallengerDidByPubKeyResponse{ChallengerDid: &challengers}, nil
}

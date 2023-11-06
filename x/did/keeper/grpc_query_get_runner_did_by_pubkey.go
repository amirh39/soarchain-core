package keeper

import (
	"context"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRunnerDidByPubKey(goCtx context.Context, req *types.QueryGetRunnerDidByPubKeyRequest) (*types.QueryGetRunnerDidByPubKeyResponse, error) {
	if req == nil || req.Pubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "[GetRunnerDidByPubKey] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	runners, found := k.GetRunnerDidUsingPubKey(ctx, req.Pubkey)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "[GetRunnerDidByPubKey][GetRunnerDidUsingPubKey] failed. Couldn't find a valid runner.")
	}

	return &types.QueryGetRunnerDidByPubKeyResponse{RunnerDid: &runners}, nil
}

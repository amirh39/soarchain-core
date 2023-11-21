package keeper

import (
	"context"

	"github.com/soar-robotics/soarchain-core/x/did/types"

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

	runner, found := k.GetRunnerDidUsingPubKey(ctx, req.Pubkey)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[GetRunnerDidByPubKey][GetRunnerDidUsingPubKey] failed. Couldn't find a valid runner for this pubkey: [ %s ] .", req.Pubkey)
	}

	return &types.QueryGetRunnerDidByPubKeyResponse{RunnerDid: &runner}, nil
}

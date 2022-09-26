package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

func (k msgServer) GenChallenger(goCtx context.Context, msg *types.MsgGenChallenger) (*types.MsgGenChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGenChallengerResponse{}, nil
}

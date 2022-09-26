package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGenClientResponse{}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

func (k msgServer) ChallengeService(goCtx context.Context, msg *types.MsgChallengeService) (*types.MsgChallengeServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgChallengeServiceResponse{}, nil
}

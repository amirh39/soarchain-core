package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SelectRandomChallenger(goCtx context.Context, msg *types.MsgSelectRandomChallenger) (*types.MsgSelectRandomChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allChallengers := k.GetAllChallenger(ctx)
	multiplier := int(len(allChallengers))

	vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
	if vrfErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "VRF error!")
	}

	return &types.MsgSelectRandomChallengerResponse{RandomChallenger: vrfData.SelectedChallenger}, nil
}

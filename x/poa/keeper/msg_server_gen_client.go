package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, isFound := k.GetClient(ctx, msg.Address)
	_, isFoundAsChallenger := k.GetChallenger(ctx, msg.Address)
	_, isFoundAsRunner := k.GetRunner(ctx, msg.Address)

	if isFound || isFoundAsChallenger || isFoundAsRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client address is already registered.")
	}

	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	// Save client into storage
	newClient := types.Client{
		Index:              msg.Address,
		Address:            msg.Address,
		Registrant:         msg.Creator,
		Score:              strconv.FormatFloat(initialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	}

	k.SetClient(ctx, newClient)

	return &types.MsgGenClientResponse{}, nil
}

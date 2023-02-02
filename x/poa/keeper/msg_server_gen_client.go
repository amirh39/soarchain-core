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

	_, isFound := k.GetClient(ctx, msg.Pubkey)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client pubkey is already registered.")
	}

	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	// Save client into storage
	newClient := types.Client{
		Index:              msg.Pubkey,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(initialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	}

	k.SetClient(ctx, newClient)

	// Register Motus client into Motus Wallet object
	_, isFoundWallet := k.GetMotusWallet(ctx, msg.Creator)
	_, isFoundAsChallenger := k.GetChallenger(ctx, msg.Creator)
	_, isFoundAsRunner := k.GetRunner(ctx, msg.Creator)

	if isFoundWallet || isFoundAsChallenger || isFoundAsRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client address is already registered.")
	}

	newMotusWallet := types.MotusWallet{
		Index:  msg.Creator,
		Client: &newClient,
	}
	k.SetMotusWallet(ctx, newMotusWallet)

	return &types.MsgGenClientResponse{}, nil
}

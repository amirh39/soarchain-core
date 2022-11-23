package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) GenChallenger(goCtx context.Context, msg *types.MsgGenChallenger) (*types.MsgGenChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Check if exists
	_, isFound := k.GetChallenger(ctx, msg.Creator)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Challenger is already registered in storage.")
	}

	// 2. Check stake amount
	requiredStake, _ := sdk.ParseCoinsNormalized("2000soar")
	stakedAmount, _ := sdk.ParseCoinsNormalized(msg.StakeAmount)
	if stakedAmount.GetDenomByIndex(0) != "soar" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid coin denominator")
	}
	if stakedAmount.IsAllLT(requiredStake) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Staked amount is:"+stakedAmount.String()+"less than required stake amount"+requiredStake.String())
	}

	// 3. Transfer stakedAmount to contract:
	challengerAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, challengerAddr, types.ModuleName, requiredStake)

	newChallenger := types.Challenger{
		Index:        msg.Creator,
		Address:      msg.Creator,
		Score:        sdk.ZeroInt().String(),
		StakedAmount: stakedAmount.String(),
		NetEarnings:  "",
		Type:         "",
		IpAddr:       "",
	}

	k.SetChallenger(ctx, newChallenger)

	return &types.MsgGenChallengerResponse{}, nil
}

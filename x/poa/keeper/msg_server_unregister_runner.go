package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterRunner(goCtx context.Context, msg *types.MsgUnregisterRunner) (*types.MsgUnregisterRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check runner
	runner, isFoundRunner := k.GetRunner(ctx, msg.RunnerAddress)
	if !isFoundRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UnregisterRunner][GetRunner] failed. Runner is not registered.")
	}
	// Check runner belongs to msg.Creator's address
	if runner.Address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Runner doesn't belong to msg.Creator's address!")
	}

	// Query the staked amount and refund
	msgSenderAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[UnregisterRunner][AccAddressFromBech32] failed. Sender address is not valid."+err.Error())
	}

	stakedAmountStr := runner.StakedAmount
	stakedAmount, err := sdk.ParseCoinsNormalized(stakedAmountStr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterRunner][ParseCoinsNormalized] failed. Stake amount is not valid."+err.Error())
	}

	transferErr2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
	if transferErr2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterRunner][SendCoinsFromModuleToAccount] failed. couldn't send coins.")
	}

	// Remove runner
	k.RemoveRunner(ctx, msg.RunnerAddress)

	return &types.MsgUnregisterRunnerResponse{}, nil
}

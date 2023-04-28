package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterGuard(goCtx context.Context, msg *types.MsgUnregisterGuard) (*types.MsgUnregisterGuardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check guard
	guard, isFound := k.GetGuard(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UnregisterGuard][GetGuard] failed. Guard is not registered, Not authorized.")
	}

	msgSenderAddress, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[UnregisterGuard][AccAddressFromBech32] failed. Sender address is not valid.")
	}

	if guard.V2XChallenger.Address == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[UnregisterGuard] failed. V2XChallenger address is not valid.")
	}

	// check v2x challenger
	v2xChallenger, isFoundV2XChallenger := k.GetChallenger(ctx, guard.V2XChallenger.Address)
	if isFoundV2XChallenger {
		stakedAmountStr := v2xChallenger.StakedAmount
		stakedAmount, err := sdk.ParseCoinsNormalized(stakedAmountStr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterGuard][ParseCoinsNormalized] failed. couldn't parse out a list of coins!"+err.Error())
		}

		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[UnregisterGuard][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
		}
		k.RemoveChallenger(ctx, guard.V2XChallenger.Address)
	}

	// check v2n challenger
	v2nChallenger, isFoundV2NChallenger := k.GetChallenger(ctx, guard.V2NChallenger.Address)
	if isFoundV2NChallenger {
		stakedAmountStr := v2nChallenger.StakedAmount
		stakedAmount, _ := sdk.ParseCoinsNormalized(stakedAmountStr)
		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[UnregisterGuard][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
		}
		k.RemoveChallenger(ctx, guard.V2NChallenger.Address)
	}

	// check runner
	runner, isFoundRunner := k.GetRunner(ctx, guard.Runner.Address)
	if isFoundRunner {
		stakedAmountStr := runner.StakedAmount
		stakedAmount, err := sdk.ParseCoinsNormalized(stakedAmountStr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterGuard][ParseCoinsNormalized] failed. Stake amount not valid"+err.Error())
		}

		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msgSenderAddress, stakedAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[UnregisterGuard][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
		}
		k.RemoveRunner(ctx, guard.Runner.Address)
	}

	k.RemoveGuard(ctx, msg.Creator)

	return &types.MsgUnregisterGuardResponse{}, nil
}

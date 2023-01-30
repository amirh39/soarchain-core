package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) ClaimRunnerRewards(goCtx context.Context, msg *types.MsgClaimRunnerRewards) (*types.MsgClaimRunnerRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	runner, isFound := k.GetRunner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target runner is not registered in the store!")
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Withdraw amount couldn't be parsed!")
	}
	earnedAmount, err := sdk.ParseCoinsNormalized(runner.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Withdraw amount couldn't be parsed!")
	}

	if earnedAmount.IsAllLT(withdrawAmount) || !withdrawAmount.DenomsSubsetOf(earnedAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Not enough coins to claim!")
	}

	runnerAccount, _ := sdk.AccAddressFromBech32(msg.Creator)
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, runnerAccount, withdrawAmount)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
	}

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)
	netEarnings := sdk.NewCoin("soar", newNetEarnings.AmountOf("soar"))

	if newNetEarnings.IsZero() {
		netEarnings = sdk.NewCoin("soar", sdk.ZeroInt())
	}

	updatedRunner := types.Runner{
		Index:              runner.Index,
		Address:            runner.Address,
		Score:              runner.Score,
		RewardMultiplier:   runner.RewardMultiplier,
		StakedAmount:       runner.StakedAmount,
		NetEarnings:        netEarnings.String(),
		IpAddr:             runner.IpAddr,
		LastTimeChallenged: runner.LastTimeChallenged,
		CoolDownTolerance:  runner.CoolDownTolerance,
		GuardAddress:       runner.GuardAddress,
	}

	k.SetRunner(ctx, updatedRunner)

	// Update runner obj in guard
	guard, isFound := k.GetGuard(ctx, runner.GuardAddress)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Guard not found")
	}
	updateGuard := types.Guard{
		Index:         guard.Index,
		GuardId:       guard.GuardId,
		V2XChallenger: guard.V2XChallenger,
		V2NChallenger: guard.V2NChallenger,
		Runner:        &updatedRunner,
	}

	k.SetGuard(ctx, updateGuard)

	return &types.MsgClaimRunnerRewardsResponse{}, nil
}

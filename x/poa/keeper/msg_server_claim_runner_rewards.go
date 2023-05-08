package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"

	params "soarchain/app/params"
)

func (k msgServer) ClaimRunnerRewards(goCtx context.Context, msg *types.MsgClaimRunnerRewards) (*types.MsgClaimRunnerRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	runner, isFound := k.GetRunner(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[ClaimRunnerRewards][GetRunner] failed. Target runner is not registered in the store by this address: [ %T ]. Make sure the address is valid and not empty.", msg.Creator)
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[ClaimRunnerRewards][ParseCoinsNormalized] failed. Withdraw amount: [ %T ] couldn't be parsed. Error: [ %T ]", msg.Amount, err)
	}

	earnedAmount, err := sdk.ParseCoinsNormalized(runner.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[ClaimRunnerRewards][ParseCoinsNormalized] failed. Withdraw amount: [ %T ] couldn't be parsed. Error: [ %T ]", runner.NetEarnings, err)
	}

	if earnedAmount.IsAllLT(withdrawAmount) || !withdrawAmount.DenomsSubsetOf(earnedAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "[ClaimRunnerRewards][IsAllLT][DenomsSubsetOf] failed. Not enough coins to claim.")
	}

	runnerAccount, _ := sdk.AccAddressFromBech32(msg.Creator)
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, runnerAccount, withdrawAmount)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ClaimRunnerRewards][IsAllLT][DenomsSubsetOf] failed. Couldn't send coins.")
	}

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)
	netEarnings := sdk.NewCoin(params.BondDenom, newNetEarnings.AmountOf(params.BondDenom))

	if newNetEarnings.IsZero() {
		netEarnings = sdk.NewCoin(params.BondDenom, sdk.ZeroInt())
	}

	updatedRunner := types.Runner{
		PubKey:             runner.PubKey,
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[ClaimRunnerRewards][GetGuard] failed. Guard not found, got: [ %T ].", runner.GuardAddress)
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

package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"

	params "soarchain/app/params"
)

func (k msgServer) ClaimRunnerRewards(goCtx context.Context, msg *types.MsgClaimRunnerRewards) (*types.MsgClaimRunnerRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Claim Runner Rewards Transaction Started ##############")

	reputation, isFound := k.GetReputationsByAddress(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[ClaimRunnerRewards][GetReputationByClientAddress] failed. Target reputation is not registered in the store by this address: [ %T ]. Make sure the address is valid and not empty.", msg.Creator)
	}

	if logger != nil {
		logger.Info("Fetching reputation from the store successfully done.", "transaction", "ClaimRunnerRewards")
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[ClaimRunnerRewards][ParseCoinsNormalized] failed. Withdraw amount: [ %T ] couldn't be parsed. Error: [ %T ]", msg.Amount, err)
	}

	earnedAmount, err := sdk.ParseCoinsNormalized(reputation.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[ClaimRunnerRewards][ParseCoinsNormalized] failed. Withdraw amount: [ %T ] couldn't be parsed. Error: [ %T ]", reputation.NetEarnings, err)
	}

	if earnedAmount.IsAllLT(withdrawAmount) || !withdrawAmount.DenomsSubsetOf(earnedAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "[ClaimRunnerRewards][DenomsSubsetOf] failed. Not enough coins to claim.")
	}

	runnerAccount, _ := sdk.AccAddressFromBech32(msg.Creator)
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, runnerAccount, withdrawAmount)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimRunnerRewards][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
	}

	if logger != nil {
		logger.Info("Transfering coins to the runner successfully done.", "transaction", "ClaimRunnerRewards")
	}

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)
	netEarnings := sdk.NewCoin(params.BondDenom, newNetEarnings.AmountOf(params.BondDenom))

	if newNetEarnings.IsZero() {
		netEarnings = sdk.NewCoin(params.BondDenom, sdk.ZeroInt())
	}

	if logger != nil {
		logger.Info("Calculating new net earning successfully done.", "transaction", "ClaimRunnerRewards")
	}

	updateReputation := types.Reputation{
		PubKey:             reputation.PubKey,
		Address:            reputation.Address,
		Score:              reputation.Score,
		RewardMultiplier:   reputation.RewardMultiplier,
		LastTimeChallenged: reputation.RewardMultiplier,
		CoolDownTolerance:  reputation.CoolDownTolerance,
		Type:               reputation.Type,
		StakedAmount:       reputation.StakedAmount,
		NetEarnings:        netEarnings.String(),
	}
	k.SetReputation(ctx, updateReputation)

	if logger != nil {
		logger.Info("Updating target reputation successfully done.", "transaction", "ClaimRunnerRewards")
	}

	log.Println("############## End of Claim Runner Rewards Transaction ##############")

	return &types.MsgClaimRunnerRewardsResponse{}, nil
}

package keeper

import (
	"context"
	"log"

	params "github.com/amirh39/soarchain-core/app/params"
	"github.com/amirh39/soarchain-core/x/did/constants"
	"github.com/amirh39/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimChallengerRewards(goCtx context.Context, msg *types.MsgClaimChallengerRewards) (*types.MsgClaimChallengerRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Claim Challenger Rewards Transaction Started ##############")

	reputation, isFound := k.GetReputationsByAddress(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[ClaimChallengerRewards][GetReputationsByAddress] failed. Target challenger is not registered in the store by this address: [ %T ]. Make sure the address is valid and not empty.", msg.Creator)
	}
	if reputation.Type != constants.V2NChallengerType && reputation.Type != constants.V2XChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[ClaimChallengerRewards] failed. The creator is not registered as a valid challenger type.")
	}

	if logger != nil {
		logger.Info("Fetching challenger from the store successfully done.", "transaction", "ClaimChallengerRewards")
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[ClaimChallengerRewards][ParseCoinsNormalized] failed. Withdraw amount: [ %T ] couldn't be parsed. Error: [ %T ]", msg.Amount, err)
	}

	earnedAmount, err := sdk.ParseCoinsNormalized(reputation.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[ClaimChallengerRewards][ParseCoinsNormalized] failed. Withdraw amount: [ %T ] couldn't be parsed. Error: [ %T ]", reputation.NetEarnings, err)
	}
	if earnedAmount == nil || withdrawAmount == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimChallengerRewards] failed. Failed to retrieve either earned amount or withdrawal amount.")
	}
	if earnedAmount.IsAllLT(withdrawAmount) || !withdrawAmount.DenomsSubsetOf(earnedAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "[ClaimChallengerRewards] failed. Claimed amount exceeds the earned amount or is not a subset of the earned amount.")
	}

	challengerAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, challengerAddress, withdrawAmount)

	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimChallengerRewards][SendCoinsFromModuleToAccount] failed. Couldn't send coins from PoA module to the challenger account.")
	}

	if logger != nil {
		logger.Info("Transferring coins to the challenger account successfully done.", "transaction", "ClaimChallengerRewards")
	}

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)
	netEarnings := sdk.NewCoin(params.BondDenom, newNetEarnings.AmountOf(params.BondDenom))

	if logger != nil {
		logger.Info("Calculating new net earning successfully done.", "transaction", "ClaimChallengerRewards")
	}

	updateReputation := types.Reputation{
		PubKey:             reputation.PubKey,
		Address:            reputation.Address,
		Score:              reputation.Score,
		RewardMultiplier:   reputation.RewardMultiplier,
		LastTimeChallenged: reputation.LastTimeChallenged,
		CoolDownTolerance:  reputation.CoolDownTolerance,
		Type:               reputation.Type,
		StakedAmount:       reputation.StakedAmount,
		NetEarnings:        netEarnings.String(),
	}
	k.SetReputation(ctx, updateReputation)

	if logger != nil {
		logger.Info("Updating target challenger successfully done.", "transaction", "ClaimChallengerRewards")
	}

	log.Println("############## End of Claim Challenger Rewards Transaction ##############")

	return &types.MsgClaimChallengerRewardsResponse{}, nil
}

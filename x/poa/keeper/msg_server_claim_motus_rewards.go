package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	params "soarchain/app/params"
	"soarchain/x/poa/types"

	didtypes "soarchain/x/did/types"
)

func (k msgServer) ClaimMotusRewards(goCtx context.Context, msg *types.MsgClaimMotusRewards) (*types.MsgClaimMotusRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Claim Motus Rewards Transaction Started ##############")

	reputation, isFound := k.didKeeper.GetReputationByClientAddress(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimMotusRewards][GetReputation] failed. Creator is not valid address.")
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimMotusRewards][ParseCoinsNormalized] failed. Couldn't parse withdrawal amount.")
	}

	earnedAmount, err := sdk.ParseCoinsNormalized(reputation.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimMotusRewards][ParseCoinsNormalized] failed. Couldn't parse withdrawal amount.")
	}

	if earnedAmount.IsAllLT(withdrawAmount) || !withdrawAmount.DenomsSubsetOf(earnedAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "[ClaimMotusRewards][IsAllLT][DenomsSubsetOf] failed. Not enough coins to claim.")
	}

	clientAccount, _ := sdk.AccAddressFromBech32(msg.Creator)
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, withdrawAmount)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimMotusRewards][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
	}

	if logger != nil {
		logger.Info("Transfering coins to the target client successfully done.", "transaction", "ClaimMotusRewards")
	}

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)
	netEarnings := sdk.NewCoin(params.BondDenom, newNetEarnings.AmountOf(params.BondDenom))

	if newNetEarnings.IsZero() {
		netEarnings = sdk.NewCoin(params.BondDenom, sdk.ZeroInt())
	}

	if logger != nil {
		logger.Info("Calculating new net earning successfully done.", "transaction", "ClaimMotusRewards")
	}

	updatedReputation := didtypes.Reputation{
		Index:       reputation.Index,
		NetEarnings: netEarnings.String(),
	}
	k.didKeeper.SetReputation(ctx, updatedReputation)

	log.Println("############## End of Claim Motus Rewards Transaction ##############")

	return &types.MsgClaimMotusRewardsResponse{}, nil
}

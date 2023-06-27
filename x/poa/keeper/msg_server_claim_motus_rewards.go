package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	params "soarchain/app/params"
	"soarchain/x/poa/types"
)

func (k msgServer) ClaimMotusRewards(goCtx context.Context, msg *types.MsgClaimMotusRewards) (*types.MsgClaimMotusRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	motusWallet, isFound := k.GetMotusWallet(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[ClaimMotusRewards][GetMotusWallet] failed. Target client is not registered in the store.")
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ClaimMotusRewards][ParseCoinsNormalized] failed. Couldn't parse withdrawal amount.")
	}

	earnedAmount, err := sdk.ParseCoinsNormalized(motusWallet.Client.NetEarnings)
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

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)
	netEarnings := sdk.NewCoin(params.BondDenom, newNetEarnings.AmountOf(params.BondDenom))

	if newNetEarnings.IsZero() {
		netEarnings = sdk.NewCoin(params.BondDenom, sdk.ZeroInt())
	}

	updatedClient := types.Client{
		Index:              motusWallet.Client.Index,
		Address:            motusWallet.Client.Address,
		Score:              motusWallet.Client.Score,
		RewardMultiplier:   motusWallet.Client.RewardMultiplier,
		NetEarnings:        netEarnings.String(),
		LastTimeChallenged: motusWallet.Client.LastTimeChallenged,
		CoolDownTolerance:  motusWallet.Client.CoolDownTolerance,
		Type:               motusWallet.Client.Type,
	}

	k.SetClient(ctx, updatedClient)

	// Update Motus wallet
	newMotusWallet := types.MotusWallet{
		Index:  motusWallet.Index,
		Client: &updatedClient,
	}
	k.SetMotusWallet(ctx, newMotusWallet)

	return &types.MsgClaimMotusRewardsResponse{}, nil
}

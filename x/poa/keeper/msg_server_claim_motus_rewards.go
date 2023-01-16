package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) ClaimMotusRewards(goCtx context.Context, msg *types.MsgClaimMotusRewards) (*types.MsgClaimMotusRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	client, isFound := k.GetClient(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target client is not registered in the store!")
	}

	withdrawAmount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Withdraw amount couldn't be parsed!")
	}

	earnedAmount, err := sdk.ParseCoinsNormalized(client.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Withdraw amount couldn't be parsed!")
	}

	if earnedAmount.IsAllLT(withdrawAmount) || !withdrawAmount.DenomsSubsetOf(earnedAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Not enough coins to claim!")
	}

	// Mint coins to module, then send it to claimer
	// ToDo: supply check will be added
	clientAccount, _ := sdk.AccAddressFromBech32(msg.Creator)

	mintErr := k.bankKeeper.MintCoins(ctx, types.ModuleName, withdrawAmount)
	if mintErr != nil {
		return nil, mintErr
	}
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, withdrawAmount)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
	}

	// Calculate new net earnings
	newNetEarnings := earnedAmount.Sub(withdrawAmount)

	updatedClient := types.Client{
		Index:              client.Index,
		Address:            client.Address,
		Registrant:         client.Registrant,
		Score:              client.Score,
		RewardMultiplier:   client.RewardMultiplier,
		NetEarnings:        newNetEarnings.String(),
		LastTimeChallenged: client.LastTimeChallenged,
		CoolDownTolerance:  client.CoolDownTolerance,
	}

	k.SetClient(ctx, updatedClient)

	return &types.MsgClaimMotusRewardsResponse{}, nil
}
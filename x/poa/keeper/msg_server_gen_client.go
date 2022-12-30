package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, isFound := k.GetClient(ctx, msg.Address)
	_, isFoundAsChallenger := k.GetChallenger(ctx, msg.Address)
	_, isFoundAsRunner := k.GetRunner(ctx, msg.Address)

	if isFound || isFoundAsChallenger || isFoundAsRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client address is already registered.")
	}

	// Registration fee
	registrationFee, _ := sdk.ParseCoinsNormalized("25000000soar")
	msgFee, err := sdk.ParseCoinsNormalized(msg.Fee)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
	}
	if msgFee.IsAllLT(registrationFee) || !msgFee.DenomsSubsetOf(registrationFee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for registration.")
	}

	// Transfer fee to the protocol, then burn it
	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, registrationFee)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins from account to POA module!")
	}
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, registrationFee)

	//
	clientAddr, _ := sdk.AccAddressFromBech32(msg.Address)

	// Save client into storage
	newClient := types.Client{
		Index:              clientAddr.String(),
		Address:            clientAddr.String(),
		Score:              sdk.NewInt(50).String(), // Base Score
		RewardMultiplier:   sdk.ZeroInt().String(),
		NetEarnings:        sdk.ZeroInt().String(),
		LastTimeChallenged: sdk.ZeroInt().String(),
	}

	k.SetClient(ctx, newClient)

	return &types.MsgGenClientResponse{}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterClient(goCtx context.Context, msg *types.MsgUnregisterClient) (*types.MsgUnregisterClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the client exists
	client, isFound := k.GetClient(ctx, msg.Pubkey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UnregisterClient][GetClient] failed. Client is not registered.")
	}

	// Check if authorized
	if client.Address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[UnregisterClient] failed. Registrant is not recognized.")
	}

	// Get Motus Wallet
	motusWallet, isFoundMotusWallet := k.GetMotusWallet(ctx, client.Address)
	if !isFoundMotusWallet {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UnregisterClient][GetMotusWallet] failed. Motus wallet is not registered.")
	}

	// Transfer claimmable rewards
	earnedAmount, err := sdk.ParseCoinsNormalized(client.NetEarnings)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterClient][ParseCoinsNormalized] failed. Withdraw amount couldn't be parsed."+err.Error())
	}
	clientAccount, _ := sdk.AccAddressFromBech32(client.Address)
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, earnedAmount)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[UnregisterClient][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
	}

	// Remove client
	k.RemoveClient(ctx, msg.Pubkey)

	// Remove motus wallet
	k.RemoveMotusWallet(ctx, motusWallet.Index)

	return &types.MsgUnregisterClientResponse{}, nil
}

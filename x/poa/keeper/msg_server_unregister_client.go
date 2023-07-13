package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) UnregisterClient(goCtx context.Context, msg *types.MsgUnregisterClient) (*types.MsgUnregisterClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Unregister client Transaction Started ##############")

	// Check if the client exists
	client, isFound := k.GetClient(ctx, msg.Pubkey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UnregisterClient][GetClient] failed. Client is not registered.")
	}

	// Check if authorized
	if client.Address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[UnregisterClient] failed. Registrant is not recognized.")
	}

	if logger != nil {
		logger.Info("Authorizing client successfully done.", "transaction", "UnregisterClient")
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

	if logger != nil {
		logger.Info("Transfering reward successfully done.", "transaction", "UnregisterClient")
	}

	// Remove client
	k.RemoveClient(ctx, msg.Pubkey)

	if logger != nil {
		logger.Info("Removing client successfully done.", "transaction", "UnregisterClient")
	}

	// Remove motus wallet
	k.RemoveMotusWallet(ctx, motusWallet.Index)

	if logger != nil {
		logger.Info("Removing motus wallet successfully done.", "transaction", "UnregisterClient")
	}

	log.Println("############## End of Unregister client Transaction ##############")

	return &types.MsgUnregisterClientResponse{}, nil
}

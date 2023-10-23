package keeper

import (
	"context"
	"log"

	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeactivateDid(goCtx context.Context, msg *types.MsgDeactivateDid) (*types.MsgDeactivateDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	log.Println("############## Deactivating a did Transaction is Started ##############")

	if msg.FromAddress == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][Validate address] failed. make sure using valid address.")
	}

	_, found := k.GetClientDid(ctx, msg.FromAddress)
	if found {
		k.RemoveClientDid(ctx, msg.FromAddress)

		log.Println("############## End of Deactivating did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	_, found = k.GetRunnerDid(ctx, msg.FromAddress)
	if found {
		k.RemoveRunnerDid(ctx, msg.FromAddress)

		log.Println("############## End of Deactivating did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	_, found = k.GetChallengerDid(ctx, msg.FromAddress)
	if found {
		k.RemoveChallengerDid(ctx, msg.FromAddress)

		log.Println("############## End of Deactivating did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid] failed. Could not find valid did.")
}

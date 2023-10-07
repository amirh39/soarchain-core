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

	_, found := k.GetClientDidDocument(ctx, msg.Did)
	if found {
		k.RemoveClientDid(ctx, msg.Did)

		log.Println("############## End of Deactivating did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	_, found = k.GetRunnerDidDocument(ctx, msg.Did)
	if found {
		k.RemoveRunnerDid(ctx, msg.Did)

		log.Println("############## End of Deactivating did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	_, found = k.GetChallengerDidDocument(ctx, msg.Did)
	if found {
		k.RemoveChallengerDid(ctx, msg.Did)

		log.Println("############## End of Deactivating did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid] failed. Could not find valid did.")
}

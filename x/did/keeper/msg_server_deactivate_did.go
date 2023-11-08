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

	if msg.Creator == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][Validate from address] failed. make sure using valid address.")
	}

	_, found := k.GetClientDid(ctx, msg.Creator)
	if found {

		k.RemoveClientDid(ctx, msg.Creator)

		error := k.poaKeeper.RemoveClientReputation(ctx, msg.Creator)
		if error != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][RemoveClientReputation] failed. There is an error during deleteting client did. Make sure using valid client did address.")
		}

		log.Println("############## End of Deactivating Client did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	_, found = k.GetRunnerDid(ctx, msg.Creator)
	if found {

		if msg.Creator == "" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][Validate runner creator address] failed. make sure using valid address.")
		}

		k.RemoveRunnerDid(ctx, msg.Creator)

		error := k.poaKeeper.RemoveRunnerReputation(ctx, msg.Creator)
		if error != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][RemoveClientReputation] failed. There is an error during deleteting runner did. Make sure using valid runner did address.")
		}

		log.Println("############## End of Deactivating Runner did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	_, found = k.GetChallengerDid(ctx, msg.Creator)
	if found {

		if msg.Creator == "" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][Validate challenger creator address] failed. make sure using valid address.")
		}

		k.RemoveChallengerDid(ctx, msg.Creator)

		error := k.poaKeeper.RemoveChallengerReputation(ctx, msg.Creator)
		if error != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][RemoveClientReputation] failed. There is an error during deleteting runner did. Make sure using valid runner did address.")
		}

		log.Println("############## End of Deactivating Challenger did Transaction ##############")

		return &types.MsgDeactivateDidResponse{}, nil
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid] failed. Could not find valid did.")
}

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
	logger := k.Logger(ctx)

	log.Println("############## Deactivating a did Transaction is Started ##############")

	documentWithSequence, found := k.GetDidDocument(ctx, msg.Did)
	doc := documentWithSequence.Document
	if !found || documentWithSequence.Empty() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][GetDidDocument] failed. Did is not registered.")
	}
	if documentWithSequence.Deactivated() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][GetDidDocument] failed. Did was already deactivated.")
	}

	if logger != nil {
		logger.Info("Checking for valid did successfully done.", "transaction", "DeactivateDid", "DocumentWithSequence", documentWithSequence)
	}

	newSequence, err := k.VerifyDidOwnership(doc, documentWithSequence.Sequence, documentWithSequence.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DeactivateDid][VerifyDidOwnership] failed. Did not belong to the creator.")
	}

	k.SetDidDocument(ctx, msg.Did, documentWithSequence.Deactivate(newSequence))

	log.Println("############## End of Deactivating did Transaction ##############")

	return &types.MsgDeactivateDidResponse{}, nil

}

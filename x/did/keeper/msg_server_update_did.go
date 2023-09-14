package keeper

import (
	"context"
	"log"

	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateDid(goCtx context.Context, msg *types.MsgUpdateDid) (*types.MsgUpdateDidResponse, error) {

	log.Println("############## Updating a did Transaction Started ##############")

	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	documentWithSeq, found := k.GetDidDocumentWithSequence(ctx, msg.Did)
	if documentWithSeq.Empty() || !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDid][GetDidDocument] failed. Did is not registered.")
	}
	if documentWithSeq.Deactivated() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDid][GetDidDocument] failed. Did was already deactivated.")
	}

	if logger != nil {
		logger.Info("Checking for valid did successfully done.", "transaction", "UpdateDid", "DocumentWithSequence", documentWithSeq)
	}

	newSeq, err := k.VerifyDidOwnership(msg.Document, documentWithSeq.Sequence, documentWithSeq.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[UpdateDid][VerifyDidOwnership] failed. Did not belong to the creator.")
	}

	if logger != nil {
		logger.Info("Verifying ownership of did successfully done.", "transaction", "UpdateDid", "DocumentWithSequence", documentWithSeq)
	}

	newDocWithSeq := types.NewDidDocumentWithSeq(msg.Document, newSeq)
	k.SetDidDocument(ctx, msg.Did, newDocWithSeq)

	log.Println("############## End of Updating a did Transaction ##############")

	return &types.MsgUpdateDidResponse{}, nil

}

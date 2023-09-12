package keeper

import (
	"context"
	"log"

	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenDid(goCtx context.Context, msg *types.MsgGenDid) (*types.MsgGenDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a did Transaction Started ##############")

	documentWithSequence, found := k.GetDidDocument(ctx, msg.Did)
	if found || !documentWithSequence.Empty() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDid][GetDidDocument] failed. Did is already registered.")
	}

	seq := types.InitialSequence
	if logger != nil {
		logger.Info("VerifyDidOwnership", "transaction", "GenDid", "msg.Document", msg.Document)
	}
	_, err := k.VerifyDidOwnership(msg.Document, seq, msg.Document, msg.VerificationMethodId, msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDid][VerifyDidOwnership] failed. Did not belong to the creator.")
	}

	didDocument := types.NewDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetDidDocument(ctx, msg.Did, didDocument)

	if logger != nil {
		logger.Info("Generating did successfully done.", "transaction", "GenDid", "document", didDocument)
	}

	log.Println("############## End of Generating did Transaction ##############")

	return &types.MsgGenDidResponse{}, nil
}

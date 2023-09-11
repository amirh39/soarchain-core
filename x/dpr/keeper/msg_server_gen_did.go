package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"
)

func (k msgServer) GenDpr(goCtx context.Context, msg *types.MsgGenDpr) (*types.MsgGenDprResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//logger := k.Logger(ctx)

	log.Println("############## Generating a dpr Transaction Started ##############")

	//documentWithSequence, found := k.GenDpr(ctx, msg.Dpr)
	// if found || !documentWithSequence.Empty() {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDid][GetDidDocument] failed. Did is already registered.")
	// }

	// seq := types.InitialSequence
	// if logger != nil {
	// 	logger.Info("VerifyDidOwnership", "transaction", "GenDid", "msg.Document", msg.Document)
	// }
	// _, err := k.VerifyDidOwnership(msg.Document, seq, msg.Document, msg.VerificationMethodId, msg.Signature)
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDid][VerifyDidOwnership] failed. Did not belong to the creator.")
	// }

	// didDocument := types.NewDidDocumentWithSeq(msg.Document, uint64(seq))
	// k.SetDidDocument(ctx, msg.Did, didDocument)

	// if logger != nil {
	// 	logger.Info("Generating did successfully done.", "transaction", "GenDid", "document", didDocument)
	// }

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgGenDprResponse{}, nil
}

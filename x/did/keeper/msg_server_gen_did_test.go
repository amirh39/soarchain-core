package keeper_test

import (
	"fmt"
	k "soarchain/x/did/keeper"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Gen_Did() {

	helper.Run("TestGenDid", func() {
		helper.Setup()
		keeper := helper.App.DidKeeper

		ctx := sdk.WrapSDKContext(helper.Ctx)
		msgServer := k.NewMsgServerImpl(keeper)

		documentWithSequence, privKey := NewDIDDocumentWithSeq(Did)
		doc := documentWithSequence.Document
		sig, error := types.Sign(doc, documentWithSequence.Sequence, privKey)
		helper.Require().NoError(error)

		res, err := msgServer.GenDid(ctx, &types.MsgGenDid{
			Document:        documentWithSequence.Document,
			Signature:       sig,
			Certificate:     Certificate,
			ClientSignature: ClientSignature,
			Creator:         ADDRESS,
		})
		helper.Require().NoError(err)
		didDocument, found := keeper.GetDidDocument(helper.Ctx, Did)
		fmt.Print("didDocument------------------->", didDocument)
		helper.Require().Equal(found, true)
		helper.Require().NotNil(res)
		helper.Require().NoError(err)

	})
}

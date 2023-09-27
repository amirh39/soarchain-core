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

		documentWithSequence, _ := NewDIDDocumentWithSeq(Did)
		helper.Require().NotEmpty(documentWithSequence)

		res, err := msgServer.GenDid(ctx, &types.MsgGenDid{
			Document:    documentWithSequence.Document,
			Signature:   Signature,
			Certificate: Certificate,
			Creator:     ADDRESS,
		})
		helper.Require().NoError(err)
		didDocument, found := keeper.GetDidDocument(helper.Ctx, Did)
		fmt.Print("didDocument------------------->", didDocument)
		helper.Require().Equal(found, true)
		helper.Require().NotNil(res)
		helper.Require().NoError(err)

	})
}

package keeper_test

import (
	"fmt"
	k "soarchain/x/did/keeper"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poatypes "soarchain/x/poa/types"
)

func (helper *KeeperTestHelper) Test_Gen_Client() {

	helper.Run("TestGenClient", func() {
		helper.Setup()
		keeper := helper.App.DidKeeper
		poakeeper := helper.App.PoaKeeper

		ctx := sdk.WrapSDKContext(helper.Ctx)
		msgServer := k.NewMsgServerImpl(keeper)

		item := poatypes.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
			MasterAccount: MASTER_ACCOUNT,
		}
		poakeeper.SetMasterKey(helper.Ctx, item)
		updatedFactoryKeyList := poatypes.FactoryKeys{
			Id:          uint64(1),
			FactoryCert: Certificate,
		}
		poakeeper.SetFactoryKeys(helper.Ctx, updatedFactoryKeyList)
		deviceCert := poakeeper.GetAllFactoryKeys(helper.Ctx)
		helper.Require().NotNil(deviceCert)

		documentWithSequence, _ := NewDIDDocumentWithSeq(Did)
		helper.Require().NotEmpty(documentWithSequence)

		res, err := msgServer.GenClient(ctx, &types.MsgGenClient{
			Document:    documentWithSequence.Document,
			Signature:   Signature,
			Certificate: Certificate,
			Creator:     ADDRESS,
		})
		didDocument, found := keeper.GetClientDid(helper.Ctx, documentWithSequence.Document.Address)
		fmt.Print("didDocument------------------->", didDocument)
		helper.Require().Equal(found, true)
		if err != nil {
			helper.Require().NotNil(err)
		} else {
			helper.Require().NotNil(res)
			helper.Require().NoError(err)
		}
	})
}
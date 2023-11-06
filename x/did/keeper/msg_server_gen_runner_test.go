package keeper_test

import (
	"fmt"
	k "soarchain/x/did/keeper"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poatypes "soarchain/x/poa/types"
)

func (helper *KeeperTestHelper) Test_Gen_Runner() {

	helper.Run("TestGenRunner", func() {
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

		documentWithSequence, _ := NewRunnerDidDocumentWithSeq(Did)
		helper.Require().NotEmpty(documentWithSequence)

		res, err := msgServer.GenRunner(ctx, &types.MsgGenRunner{
			Document:    documentWithSequence.Document,
			Signature:   Signature,
			Certificate: Certificate,
			Creator:     ADDRESS,
			RunnerStake: RunnerStake,
			RunnerIp:    RunnerIp,
		})
		didDocument, found := keeper.GetRunnerDid(helper.Ctx, documentWithSequence.Document.Address)
		fmt.Print("didDocument------------------->", didDocument)
		if err != nil {
			helper.Require().NotNil(err)
			helper.Require().Equal(found, false)
		} else {
			helper.Require().NotNil(res)
			helper.Require().NoError(err)
			helper.Require().Equal(found, true)
		}
	})
}

package keeper_test

import (
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
		newDid := types.RunnerDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Address: ADDRESS,
		}

		res, err := msgServer.GenRunner(ctx, &types.MsgGenRunner{
			Document:    &newDid,
			Signature:   Signature,
			Certificate: Certificate,
			Creator:     ADDRESS,
			RunnerStake: RunnerStake,
		})
		if err != nil {
			helper.Require().NotNil(err)
			helper.Require().Nil(res)
		} else {
			didDocument, found := keeper.GetRunnerDid(helper.Ctx, newDid.Address)
			helper.Require().NotNil(didDocument)
			helper.Require().NoError(err)
			helper.Require().Equal(found, true)
		}
	})
}

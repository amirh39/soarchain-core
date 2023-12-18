package keeper_test

import (
	k "github.com/amirh39/soarchain-core/x/did/keeper"
	"github.com/amirh39/soarchain-core/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poaTypes "github.com/amirh39/soarchain-core/x/poa/types"
)

func (helper *KeeperTestHelper) Test_Gen_Client() {

	helper.Run("TestGenClient", func() {
		helper.Setup()
		keeper := helper.App.DidKeeper
		poaKeeper := helper.App.PoaKeeper

		ctx := sdk.WrapSDKContext(helper.Ctx)
		msgServer := k.NewMsgServerImpl(keeper)

		item := poaTypes.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
			MasterAccount: MASTER_ACCOUNT,
		}
		poaKeeper.SetMasterKey(helper.Ctx, item)

		factoryCert1 := poaTypes.FactoryKeys{
			Id:          uint64(0),
			FactoryCert: FactoryCert,
		}
		factoryCert2 := poaTypes.FactoryKeys{
			Id:          uint64(1),
			FactoryCert: FactoryCert2,
		}
		poaKeeper.SetFactoryKeys(helper.Ctx, factoryCert1)
		poaKeeper.SetFactoryKeys(helper.Ctx, factoryCert2)

		factoryCerts := poaKeeper.GetAllFactoryKeys(helper.Ctx)
		helper.Require().NotEmpty(factoryCerts)

		res, err := msgServer.GenClient(ctx, &types.MsgGenClient{
			Signature:   Signature,
			Certificate: Certificate,
			Creator:     ADDRESS2,
		})
		helper.Require().NotNil(res)
		helper.Require().NoError(err)

		didDocument, found := keeper.GetClientDid(helper.Ctx, ADDRESS2)
		helper.Require().NotNil(didDocument)
		helper.Require().Equal(found, true)
	})
}

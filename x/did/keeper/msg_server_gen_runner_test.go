package keeper_test

import (
	"log"
	"soarchain/app/params"
	k "soarchain/x/did/keeper"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poaTypes "soarchain/x/poa/types"

	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (helper *KeeperTestHelper) Test_Gen_Runner() {

	helper.Run("TestGenRunner", func() {
		helper.Setup()
		keeper := helper.App.DidKeeper
		poaKeeper := helper.App.PoaKeeper
		bankKeeper := helper.App.BankKeeper
		accountKeeper := helper.App.AccountKeeper

		addr := sdk.MustAccAddressFromBech32(ADDRESS2)
		actorAmount := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(MintAmount)))

		accountKeeper.SetAccount(helper.Ctx, authTypes.NewBaseAccountWithAddress(addr))

		bankKeeper.MintCoins(helper.Ctx, types.ModuleName, actorAmount)
		bankKeeper.SendCoinsFromModuleToAccount(helper.Ctx, types.ModuleName, addr, actorAmount)

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

		deviceCert := poaKeeper.GetAllFactoryKeys(helper.Ctx)
		helper.Require().NotNil(deviceCert)

		_, err := msgServer.GenRunner(ctx, &types.MsgGenRunner{
			Signature:   Signature,
			Certificate: Certificate,
			Creator:     ADDRESS2,
			RunnerStake: RunnerStake,
		})

		helper.Require().Nil(err)

		didDocument, found := keeper.GetRunnerDid(helper.Ctx, ADDRESS2)
		helper.Require().NotNil(didDocument)

		helper.Require().NoError(err)
		helper.Require().Equal(found, true)
		log.Println(didDocument)

	})
}

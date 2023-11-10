package keeper_test

import (
	"soarchain/app/params"
	k "soarchain/x/did/keeper"
	"soarchain/x/did/types"
	poaTypes "soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (helper *KeeperTestHelper) Test_Gen_Challenger() {
	helper.Run("TestGenChallenger", func() {
		helper.Setup()
		keeper := helper.App.DidKeeper
		poaKeeper := helper.App.PoaKeeper
		bankKeeper := helper.App.BankKeeper
		accountKeeper := helper.App.AccountKeeper

		ctx := sdk.WrapSDKContext(helper.Ctx)
		msgServer := k.NewMsgServerImpl(keeper)

		masterCert := poaTypes.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
			MasterAccount: MASTER_ACCOUNT,
		}
		poaKeeper.SetMasterKey(helper.Ctx, masterCert)

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

		amount := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(1000000000000000000)))
		bankKeeper.MintCoins(helper.Ctx, types.ModuleName, amount)

		addr := sdk.MustAccAddressFromBech32(ADDRESS2)
		accountKeeper.SetAccount(helper.Ctx, authtypes.NewBaseAccountWithAddress(addr))
		bankKeeper.SendCoinsFromModuleToAccount(helper.Ctx, types.ModuleName, addr, amount)

		res, err := msgServer.GenChallenger(ctx, &types.MsgGenChallenger{
			Signature:       Signature,
			Certificate:     Certificate,
			Creator:         ADDRESS2,
			ChallengerStake: Challenger_StakedAmount,
			ChallengerType:  Challenger_Type,
		})

		helper.Require().Nil(err)
		helper.Require().NotNil(res)

	})
}

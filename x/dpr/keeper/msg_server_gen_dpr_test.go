package keeper_test

import (
	"log"
	"soarchain/app/params"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "soarchain/x/did/types"
)

func (helper *KeeperTestHelper) Test_Gen_DPR() {

	helper.Run("TestGenDpr", func() {
		helper.Setup()
		//actor := RandomAccountAddress()
		didKeeper := helper.App.DidKeeper
		epochKeeper := helper.App.EpochKeeper
		accountKeeper := helper.App.AccountKeeper
		bankKeeper := helper.App.BankKeeper

		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		epochData := CreateEpochData(&epochKeeper, helper.Ctx)
		epochKeeper.SetEpochData(helper.Ctx, epochData)

		newDid := didtypes.ClientDid{
			Id:            Did,
			PubKey:        PUBKEY,
			SupportedPIDs: "FFFFFFF",
		}

		didDocument := didtypes.ClientDidWithSeq{
			Document: &newDid,
			Sequence: 0,
		}
		didKeeper.SetClientDid(helper.Ctx, *didDocument.Document)
		//addr, err := sdk.AccAddressFromBech32(CREATOR)

		//testAcc := helper.TestAccs[0]

		actorAmount := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(1000000000000000)))
		// testAccountPubkey := secp256k1.GenPrivKeyFromSecret([]byte("acc")).PubKey()
		// testAccountAddress := sdk.AccAddress(testAccountPubkey.Address())
		// modAcc := authtypes.NewModuleAccount(authtypes.NewBaseAccount(testAccountAddress, testAccountPubkey, 1, 0),
		// 	"mint",
		// 	"permission",
		// )
		// accountKeeper.SetModuleAccount(helper.Ctx, modAcc)

		dprModuleAcc := accountKeeper.GetModuleAddress(types.ModuleName)
		mintModuleAcc := accountKeeper.GetModuleAddress("soarmint")
		log.Println(dprModuleAcc, mintModuleAcc)
		simapp.FundAccount(helper.App.BankKeeper, helper.Ctx, sdk.AccAddress(CREATOR), actorAmount)

		log.Println(accountKeeper.GetAccount(helper.Ctx, sdk.AccAddress(CREATOR)))

		log.Println(bankKeeper.GetBalance(helper.Ctx, sdk.AccAddress(CREATOR), params.BondDenom))

		res, err := helper.MsgServer.GenDpr(ctx, &types.MsgGenDpr{
			Creator:        CREATOR,
			SupportedPIDs:  "BE1FA813",
			Duration:       45,
			DprBudget:      "1000udmotus",
			MaxClientCount: 10,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

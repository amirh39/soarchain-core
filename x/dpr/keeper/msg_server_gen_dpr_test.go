package keeper_test

import (
	"log"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "soarchain/x/did/types"
)

func (helper *KeeperTestHelper) Test_Gen_DPR() {

	helper.Run("TestGenDpr", func() {
		helper.Setup()

		didKeeper := helper.App.DidKeeper
		epochKeeper := helper.App.EpochKeeper
		accountKeeper := helper.App.AccountKeeper

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
		acc := accountKeeper.GetAccount(helper.Ctx, accountKeeper.GetModuleAddress("dpr"))
		log.Println("ACC=", acc)
		val := accountKeeper.NewAccountWithAddress(helper.Ctx, sdk.AccAddress(CREATOR))
		log.Println(sdk.AccAddress(CREATOR), "val=", val)

		has := accountKeeper.GetAllAccounts(helper.Ctx)
		log.Println(has)

		log.Println(has)

		//accountKeeper.Account(ctx)

		res, err := helper.MsgServer.GenDpr(ctx, &types.MsgGenDpr{
			Creator:       CREATOR,
			SupportedPIDs: "BE1FA813",
			Duration:      45,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

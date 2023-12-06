package keeper_test

import (
	"log"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Update_DPR() {

	helper.Run("TestUpdateDpr", func() {
		helper.Setup()

		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupActiveDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		log.Println(dpr)

		res, err := helper.MsgServer.UpdateDpr(ctx, &types.MsgUpdateDpr{
			DprId:          DprID,
			Duration:       45,
			Sender:         CREATOR,
			DprBudget:      "10udmotus",
			MaxClientCount: 1000000000,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
		dprC, _ := dprKeeper.GetDpr(helper.Ctx, DprID)
		log.Println(dprC)
	})
}

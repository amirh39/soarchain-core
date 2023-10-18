package keeper_test

import (
	"log"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Activate_DPR() {
	helper.Run("TestActivateDpr", func() {
		helper.Setup()

		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		dprFirst, _ := dprKeeper.GetDpr(helper.Ctx, dpr[0].Id)
		log.Println(dprFirst)
		res, err := helper.MsgServer.ActivateDpr(ctx, &types.MsgActivateDpr{
			Sender: CREATOR,
			DprId:  DprId,
		})
		dprSecond, _ := dprKeeper.GetDpr(helper.Ctx, dpr[0].Id)
		log.Println(dprSecond)
		helper.Require().Empty(res)
		helper.Require().Nil(err)
		helper.Require().Equal(dprSecond.IsActive, true)
		helper.Require().NotNil(dprSecond.DprEndTime)
	})
}

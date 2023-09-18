package keeper_test

import (
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Deactivate_DPR() {
	helper.Run("TestDeactivateDpr", func() {
		helper.Setup()

		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupSecondDpr(1)
		dpr[0].IsActive = true
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		res, err := helper.MsgServer.DeactivateDpr(ctx, &types.MsgDeactivateDpr{
			DprId:  DprId,
			Sender: CREATOR,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

package keeper_test

import (
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

		res, err := helper.MsgServer.ActivateDpr(ctx, &types.MsgActivateDpr{
			Sender: CREATOR,
			DprId:  DprId,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

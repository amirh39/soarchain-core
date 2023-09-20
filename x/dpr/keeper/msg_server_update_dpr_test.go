package keeper_test

import (
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

		dpr := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		res, err := helper.MsgServer.UpdateDpr(ctx, &types.MsgUpdateDpr{
			DprId:                         DprId,
			PidSupportedOneToTwnety:       true,
			PidSupportedTwentyOneToForthy: true,
			PidSupportedForthyOneToSixty:  false,
			Duration:                      45,
			Sender:                        CREATOR,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

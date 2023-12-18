package keeper_test

import (
	"github.com/amirh39/soarchain-core/x/dpr/keeper"
	"github.com/amirh39/soarchain-core/x/dpr/types"

	didtypes "github.com/amirh39/soarchain-core/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Leave_DPR() {
	helper.Run("TestLeaveDpr", func() {
		helper.Setup()

		didKeeper := helper.App.DidKeeper
		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		dprinfo := &didtypes.DprInfo{ // Note: create a pointer
			Id:      DprID,
			Claimed: "0udmotus",
		}

		newDid := didtypes.ClientDid{
			Id:       Did,
			PubKey:   PUBKEY,
			Address:  ADDRESS,
			DprInfos: []*didtypes.DprInfo{dprinfo},
		}

		didKeeper.SetClientDid(helper.Ctx, newDid)

		res, err := helper.MsgServer.LeaveDpr(ctx, &types.MsgLeaveDpr{
			Sender: ADDRESS,
			DprId:  DprID,
		})
		did, _ := didKeeper.GetClientDid(helper.Ctx, ADDRESS)

		helper.Require().Empty(res)
		helper.Require().Nil(err)
		helper.Require().Empty(did.DprInfos)

	})
}

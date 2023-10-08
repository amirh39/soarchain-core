package keeper_test

import (
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	didtypes "soarchain/x/did/types"

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

		vin := didtypes.Vehicle{
			Vin: VIN,
		}

		newDid := didtypes.ClientDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Vehicle: &vin,
		}

		didDocument := didtypes.ClientDidWithSeq{
			Document: &newDid,
			Sequence: 0,
		}
		didKeeper.SetClientDid(helper.Ctx, *didDocument.Document)
		res, err := helper.MsgServer.LeaveDpr(ctx, &types.MsgLeaveDpr{
			PubKey: PUBKEY,
			Sender: CREATOR,
			DprId:  DprId,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

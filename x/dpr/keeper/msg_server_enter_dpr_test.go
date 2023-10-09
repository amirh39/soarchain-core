package keeper_test

import (
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	didtypes "soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Enter_DPR() {
	helper.Run("TestEnterDpr", func() {
		helper.Setup()
		didKeeper := helper.App.DidKeeper
		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupDpr(1)
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
		res, err := helper.MsgServer.EnterDpr(ctx, &types.MsgEnterDpr{
			PubKey: PUBKEY,
			Sender: CREATOR,
			DprId:  DprId,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

package keeper_test

import (
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

		res, err := helper.MsgServer.GenDpr(ctx, &types.MsgGenDpr{
			Creator:       CREATOR,
			SupportedPIDs: "BE1FA813",
			Duration:      45,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

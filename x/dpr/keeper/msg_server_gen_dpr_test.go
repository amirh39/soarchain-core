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

		newDid := didtypes.DidDocument{
			Id:                            Did,
			ClientPublicKey:               PUBKEY,
			PidSupportedOneToTwnety:       true,
			PidSupportedTwentyOneToForthy: false,
			PidSupportedForthyOneToSixty:  false,
		}

		didDocument := didtypes.DidDocumentWithSeq{
			Document: &newDid,
			Sequence: 0,
		}
		didKeeper.SetDidDocument(helper.Ctx, newDid.Id, didDocument)

		res, err := helper.MsgServer.GenDpr(ctx, &types.MsgGenDpr{
			Creator:                       CREATOR,
			PidSupportedOneToTwnety:       true,
			PidSupportedTwentyOneToForthy: false,
			PidSupportedForthyOneToSixty:  false,
			Duration:                      45,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

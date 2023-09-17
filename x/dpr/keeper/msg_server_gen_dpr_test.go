package keeper_test

import (
	"fmt"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// epochKeeper "soarchain/x/epoch/keeper"
	// epochTypes "soarchain/x/epoch/types"

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
			Id:              Did,
			ClientPublicKey: PUBKEY,
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
			LengthOfDpr:                   45,
		})
		fmt.Print("res----->", res)
		fmt.Print("err----->", err)
		// helper.Require().Empty(res)
		// helper.Require().Nil(err)
	})
}

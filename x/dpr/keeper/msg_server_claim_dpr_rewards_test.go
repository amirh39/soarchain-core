package keeper_test

import (
	"log"
	didtypes "soarchain/x/did/types"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Claim_DPR() {

	helper.Run("TestClaimDpr", func() {
		helper.Setup()

		dprKeeper := helper.App.DprKeeper
		didKeeper := helper.App.DidKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		dprinfo := &didtypes.DprInfo{ // Note: create a pointer
			Id:      DprId,
			Claimed: "0",
		}

		newDid := didtypes.ClientDid{
			Id:       Did,
			PubKey:   PUBKEY,
			Address:  ADDRESS,
			DprInfos: []*didtypes.DprInfo{dprinfo},
		}

		didKeeper.SetClientDid(helper.Ctx, newDid)
		client, _ := didKeeper.GetClientDid(helper.Ctx, ADDRESS)
		log.Println(client)

		res, err := helper.MsgServer.ClaimDprRewards(ctx, &types.MsgClaimDprRewards{
			Sender: CREATOR,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
		dprC, _ := dprKeeper.GetDpr(helper.Ctx, DprID)
		log.Println(dprC)
	})
}

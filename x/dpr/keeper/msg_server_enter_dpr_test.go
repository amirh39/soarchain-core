package keeper_test

import (
	"log"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"
	poatypes "soarchain/x/poa/types"

	didtypes "soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Enter_DPR() {
	helper.Run("TestEnterDpr", func() {
		helper.Setup()
		poakeeper := helper.App.PoaKeeper
		didKeeper := helper.App.DidKeeper
		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		dprx, _ := dprKeeper.GetDpr(helper.Ctx, DprId)
		log.Println(dprx)
		helper.Require().NotEmpty(dpr)

		newDid := didtypes.ClientDid{
			Id:            Did,
			PubKey:        PUBKEY,
			SupportedPIDs: "FFFFF",
			Address:       ADDRESS,
		}

		didDocument := didtypes.ClientDidWithSeq{
			Document: &newDid,
			Sequence: 0,
		}

		reputation := poatypes.Reputation{
			PubKey:             PUBKEY,
			Address:            ADDRESS,
			Score:              ClientScroe,
			RewardMultiplier:   ClientRewardMultiplier,
			NetEarnings:        ClientNetEarnings,
			LastTimeChallenged: LastTimeChallenged,
			CoolDownTolerance:  CoolDownTolerance,
		}
		poakeeper.SetReputation(helper.Ctx, reputation)
		didKeeper.SetClientDid(helper.Ctx, *didDocument.Document)
		rep, found := poakeeper.GetReputation(helper.Ctx, PUBKEY)
		log.Println(rep, found)

		helper.MsgServer.EnterDpr(ctx, &types.MsgEnterDpr{
			Sender: ADDRESS,
			DprId:  DprId,
		})

	})
}

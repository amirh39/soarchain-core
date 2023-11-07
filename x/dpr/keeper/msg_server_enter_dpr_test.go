package keeper_test

import (
	"log"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	didtypes "soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Enter_DPR() {
	helper.Run("TestEnterDpr", func() {
		helper.Setup()
		//poakeeper := helper.App.PoaKeeper
		didKeeper := helper.App.DidKeeper
		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])

		helper.Require().NotEmpty(dpr)

		dprinfo := &didtypes.DprInfo{ // Note: create a pointer
			Id:      DprId,
			Claimed: "0udmotus",
		}

		newDid := didtypes.ClientDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Address: ADDRESS,
		}

		didDocument := didtypes.ClientDidWithSeq{
			Document: &newDid,
			Sequence: 0,
		}

		// reputation := poatypes.Reputation{
		// 	PubKey:             PUBKEY,
		// 	Address:            ADDRESS,
		// 	Score:              ClientScroe,
		// 	RewardMultiplier:   ClientRewardMultiplier,
		// 	NetEarnings:        ClientNetEarnings,
		// 	LastTimeChallenged: LastTimeChallenged,
		// 	CoolDownTolerance:  CoolDownTolerance,
		// }
		//poakeeper.SetReputation(helper.Ctx, reputation)
		didKeeper.SetClientDid(helper.Ctx, *didDocument.Document)
		// rep, found := poakeeper.GetReputation(helper.Ctx, PUBKEY)
		// log.Println(rep, found)

		res, err := helper.MsgServer.EnterDpr(ctx, &types.MsgEnterDpr{
			Sender: ADDRESS,
			DprId:  DprId,
		})

		dprx, _ := dprKeeper.GetDpr(helper.Ctx, DprId)
		helper.Require().Equal(dprx.ClientCounter, uint64(1))

		did, _ := didKeeper.GetClientDid(helper.Ctx, ADDRESS)

		helper.Require().Equal(did.DprInfos[0], dprinfo)

		helper.Require().Empty(res)
		helper.Require().Nil(err)

		test := did.GetDprInfos()

		// Create a slice to hold the IDs
		var ids []string

		// Iterate over the slice of pointers to DprInfos
		for _, infoPtr := range test {
			// Check if the pointer is not nil
			if infoPtr != nil {
				// Append the ID to the slice of IDs
				ids = append(ids, infoPtr.Id)
			}
		}

		log.Println(ids)
		log.Println(dprKeeper.GetDpr(helper.Ctx, ids[0]))
	})
}

package keeper_test

// import (
// 	"log"
// 	"soarchain/x/dpr/keeper"
// 	"soarchain/x/dpr/types"

// 	didtypes "soarchain/x/did/types"

// 	poatypes "soarchain/x/poa/types"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// func (helper *KeeperTestHelper) Test_Leave_DPR() {
// 	helper.Run("TestLeaveDpr", func() {
// 		helper.Setup()
// 		didKeeper := helper.App.DidKeeper
// 		dprKeeper := helper.App.DprKeeper
// 		poakeeper := helper.App.PoaKeeper
// 		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)

// 		ctx := sdk.WrapSDKContext(helper.Ctx)

// 		reputation := poatypes.Reputation{
// 			PubKey:             PUBKEY,
// 			Address:            ADDRESS,
// 			Score:              CREATOR,
// 			RewardMultiplier:   ClientRewardMultiplier,
// 			NetEarnings:        ClientNetEarnings,
// 			LastTimeChallenged: LastTimeChallenged,
// 			CoolDownTolerance:  CoolDownTolerance,
// 		}
// 		poakeeper.SetReputation(helper.Ctx, reputation)

// 		dpr := SetupSecondDpr(1)
// 		dprKeeper.SetDpr(helper.Ctx, dpr[0])
// 		helper.Require().NotEmpty(dpr)

// 		vin := didtypes.Vehicle{
// 			Vin: VIN,
// 		}

// 		newDid := didtypes.ClientDid{
// 			Id:      Did,
// 			PubKey:  PUBKEY,
// 			Address: ADDRESS,
// 			Vehicle: &vin,
// 		}

// 		didDocument := didtypes.ClientDidWithSeq{
// 			Document: &newDid,
// 			Sequence: 0,
// 		}
// 		didKeeper.SetClientDid(helper.Ctx, *didDocument.Document)
// 		did, found := didKeeper.GetClientDid(helper.Ctx, ADDRESS)
// 		println(found)
// 		log.Println(did)
// 		helper.MsgServer.LeaveDpr(ctx, &types.MsgLeaveDpr{
// 			PubKey: PUBKEY,
// 			Sender: CREATOR,
// 			DprId:  DprId,
// 		})
// 		// helper.Require().Empty(res)
// 		// helper.Require().Nil(err)
// 	})
// }

package keeper_test

// import (
// 	"soarchain/x/did/types"
// )

// func (helper *KeeperTestHelper) Test_DeactivateDid() {
// 	helper.Run("TestGenChallenger", func() {
// 		helper.Setup()
// 		keeper := helper.App.DidKeeper
// 		poakeeper := helper.App.PoaKeeper

// 		did, docWithSeq, _, _ := MakeTestData()
// 		keeper.SetClientDid(helper.Ctx, *docWithSeq.Document)
// 		helper.Require().NotNil(did)

// 		didDocument, found := keeper.GetClientDid(helper.Ctx, ADDRESS)
// 		helper.Require().NotNil(didDocument)
// 		helper.Require().Equal(found, true)

// 		deactivateMsg := types.MsgDeactivateDid{
// 			Did:         Did,
// 			FromAddress: ADDRESS,
// 		}
// 		helper.Require().NotNil(deactivateMsg)
// 		keeper.RemoveClientDid(helper.Ctx, ADDRESS)
// 		error := poakeeper.RemoveClientReputation(helper.Ctx, ADDRESS)
// 		if error != nil {
// 			helper.Require().NotNil(error)
// 		} else {
// 			helper.Require().Nil(error)
// 		}
// 	})
// }

package keeper_test

import (
	"soarchain/x/did/types"
)

func (helper *KeeperTestHelper) Test_DeactivateDid() {
	helper.Run("TestGenChallenger", func() {
		helper.Setup()
		keeper := helper.App.DidKeeper
		poakeeper := helper.App.PoaKeeper

		newDid := types.ClientDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Address: ADDRESS,
		}
		keeper.SetClientDid(helper.Ctx, newDid)
		helper.Require().NotNil(newDid)

		didDocument, found := keeper.GetClientDid(helper.Ctx, ADDRESS)
		helper.Require().NotNil(didDocument)
		helper.Require().Equal(found, true)

		deactivateMsg := types.MsgDeactivateDid{

			Creator: ADDRESS,
		}
		helper.Require().NotNil(deactivateMsg)
		keeper.RemoveClientDid(helper.Ctx, ADDRESS)
		error := poakeeper.RemoveClientReputation(helper.Ctx, ADDRESS)
		if error != nil {
			helper.Require().NotNil(error)
		} else {
			helper.Require().Nil(error)
		}
	})
}

package keeper_test

import (
	"log"

	"github.com/amirh39/soarchain-core/app/params"
	k "github.com/amirh39/soarchain-core/x/did/keeper"
	"github.com/amirh39/soarchain-core/x/did/types"
	poaTypes "github.com/amirh39/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_DeactivateDid() {
	helper.Run("Test_DeactivateDid", func() {
		helper.Setup()
		bankKeeper := helper.App.BankKeeper
		keeper := helper.App.DidKeeper
		poaKeeper := helper.App.PoaKeeper
		msgServer := k.NewMsgServerImpl(keeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)
		newDid := types.ClientDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Address: ADDRESS,
		}
		keeper.SetClientDid(helper.Ctx, newDid)
		helper.Require().NotNil(newDid)

		didDocument, found := keeper.GetClientDid(helper.Ctx, ADDRESS)
		rep := poaTypes.Reputation{
			PubKey:      PUBKEY,
			Address:     ADDRESS,
			NetEarnings: "10udmotus",
			Type:        "mini",
		}
		poaKeeper.SetReputation(helper.Ctx, rep)
		repx, _ := poaKeeper.GetReputationsByAddress(helper.Ctx, ADDRESS)
		log.Println(repx)
		log.Println(didDocument)
		helper.Require().NotNil(didDocument)
		helper.Require().Equal(found, true)

		amount := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(1000000000000000)))
		bankKeeper.MintCoins(helper.Ctx, "poa", amount)

		deactivateMsg := types.MsgDeactivateDid{

			Creator: ADDRESS,
		}
		res, err := msgServer.DeactivateDid(ctx, &deactivateMsg)
		_, isFound := keeper.GetClientDid(helper.Ctx, ADDRESS)
		helper.Require().Equal(isFound, false)
		helper.Require().NotNil(res)
		helper.Require().Nil(err)

	})
}

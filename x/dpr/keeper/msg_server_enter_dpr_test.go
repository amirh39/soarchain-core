package keeper_test

import (
	"log"
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

		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		vin := didtypes.Vehicle{
			Vin: VIN,
		}

		newDid := didtypes.ClientDid{
			Id:            Did,
			PubKey:        PUBKEY,
			Vehicle:       &vin,
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
			PubKey: PUBKEY,
			Sender: ADDRESS,
			DprId:  DprId,
		})

	})
}

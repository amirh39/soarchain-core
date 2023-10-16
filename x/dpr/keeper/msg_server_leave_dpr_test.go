package keeper_test

import (
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"
	poatypes "soarchain/x/poa/types"

	didtypes "soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Leave_DPR() {
	helper.Run("TestLeaveDpr", func() {
		helper.Setup()
		poaKeeper := helper.App.PoaKeeper
		didKeeper := helper.App.DidKeeper
		dprKeeper := helper.App.DprKeeper
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		dpr := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr[0])
		helper.Require().NotEmpty(dpr)

		vin := didtypes.Vehicle{
			Vin: VIN,
		}

		newDid := didtypes.ClientDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Vehicle: &vin,
			Address: ADDRESS,
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
		poaKeeper.SetReputation(helper.Ctx, reputation)
		didKeeper.SetClientDid(helper.Ctx, *didDocument.Document)

		res, err := helper.MsgServer.LeaveDpr(ctx, &types.MsgLeaveDpr{
			Sender: ADDRESS,
			DprId:  DprId,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)

	})
}

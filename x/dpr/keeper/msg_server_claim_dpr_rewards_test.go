package keeper_test

import (
	"github.com/amirh39/soarchain-core/app/params"
	didtypes "github.com/amirh39/soarchain-core/x/did/types"
	"github.com/amirh39/soarchain-core/x/dpr/keeper"
	"github.com/amirh39/soarchain-core/x/dpr/types"
	epochtypes "github.com/amirh39/soarchain-core/x/epoch/types"

	poaTypes "github.com/amirh39/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_Claim_DPR() {

	helper.Run("TestClaimDpr", func() {
		helper.Setup()
		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		bankKeeper := helper.App.BankKeeper
		dprKeeper := helper.App.DprKeeper
		didKeeper := helper.App.DidKeeper
		poaKeeper := helper.App.PoaKeeper
		epochKeeper := helper.App.EpochKeeper

		targetDpr := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, targetDpr[0])
		helper.Require().NotEmpty(targetDpr)

		mintAmount, _ := sdk.ParseCoinNormalized(targetDpr[0].DprBudget)
		bankKeeper.MintCoins(helper.Ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.Int(mintAmount.Amount))))
		mockedEpochData := epochtypes.EpochData{
			TotalEpochs: 10,
		}
		epochKeeper.SetEpochData(helper.Ctx, mockedEpochData)

		clientDid := didtypes.ClientDid{
			Id:      Did,
			PubKey:  PUBKEY,
			Address: CREATOR,
			DprInfos: []*didtypes.DprInfo{
				{
					Id:      targetDpr[0].Id,
					Claimed: "0",
				}},
		}
		didKeeper.SetClientDid(helper.Ctx, clientDid)
		clientReputation := poaTypes.Reputation{
			PubKey:      PUBKEY,
			Address:     CREATOR,
			DprEarnings: "0" + params.BondDenom,
		}
		poaKeeper.SetReputation(helper.Ctx, clientReputation)

		res, err := helper.MsgServer.ClaimDprRewards(ctx, &types.MsgClaimDprRewards{
			Sender: CREATOR,
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

package keeper_test

import (
	"fmt"

	"github.com/amirh39/soarchain-core/x/poa/types"

	k "github.com/amirh39/soarchain-core/x/poa/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_ClaimMotusReward() {

	helper.Run("TestClaimMotusReward", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		msgServer := k.NewMsgServerImpl(keeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		reputation := types.Reputation{
			PubKey:      ClientPubKey,
			Address:     ClientAddress,
			NetEarnings: "100000000udmotus",
		}
		keeper.SetReputation(helper.Ctx, reputation)
		helper.Require().NotEmpty(reputation)

		resp, err := msgServer.ClaimMotusRewards(ctx, &types.MsgClaimMotusRewards{
			Creator: ClientAddress,
			Amount:  "10udmotus",
		})
		fmt.Print("resp------>", resp)
		if err != nil {
			helper.Require().NotNil(err)
		} else {
			helper.Require().Nil(err)
		}
	})
}

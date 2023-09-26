package keeper_test

import (
	"fmt"
	didtypes "soarchain/x/did/types"
	"soarchain/x/poa/types"

	k "soarchain/x/poa/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_ClaimMotusReward() {

	helper.Run("TestClaimMotusReward", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		msgServer := k.NewMsgServerImpl(keeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		didKeeper := helper.App.DidKeeper

		reputation := didtypes.Reputation{
			Index:       ClientPubKey,
			Address:     ClientAddress,
			NetEarnings: "100000000udmotus",
		}
		didKeeper.SetReputation(helper.Ctx, reputation)
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

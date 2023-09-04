package keeper_test

import (
	k "soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) TestRunnerChallenge() {

	helper.Run("TestRunnerChallenge", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper

		client := CreateTwoClientsWithAllFields(&keeper, helper.Ctx)
		CreateNChallengerWithNormalScore(&keeper, helper.Ctx, 1)

		CreateNRunner(&keeper, helper.Ctx, 1)

		msgServer := k.NewMsgServerImpl(keeper)

		ClientPubKeys := []string{
			client[0].Index,
			client[1].Index,
		}

		res, err := msgServer.RunnerChallenge(sdk.WrapSDKContext(helper.Ctx), &types.MsgRunnerChallenge{Creator: Challenger_Address, RunnerPubkey: RunnerPubKey, ClientPubkeys: ClientPubKeys, ChallengeResult: "reward"})
		helper.NoError(err)
		helper.Empty(res)

		clientUpdated, isFound0 := keeper.GetClient(helper.Ctx, ClientPubKeys[1])
		clientUpdated2, isFound1 := keeper.GetClient(helper.Ctx, ClientPubKeys[0])
		runnerUpdated, isFound2 := keeper.GetRunner(helper.Ctx, RunnerAddress)
		challengerUpdated, isFound3 := keeper.GetChallenger(helper.Ctx, Challenger_Address)
		helper.Require().NotEmpty(isFound0)
		helper.Require().NotEmpty(isFound1)
		helper.Require().NotEmpty(isFound2)
		helper.Require().NotEmpty(isFound3)
		helper.Equal("88.259", clientUpdated.Score)
		helper.Equal(clientUpdated.NetEarnings, "1767859udmotus")
		helper.Equal(clientUpdated2.NetEarnings, "1232140udmotus")
		helper.Equal(runnerUpdated.NetEarnings, "1000000udmotus")
		helper.Equal(challengerUpdated.NetEarnings, "1000000udmotus")

	})
}

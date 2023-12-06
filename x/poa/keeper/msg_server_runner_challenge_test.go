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

		reputation := CreateTwoReputationsWithAllFields(&keeper, helper.Ctx)
		CreateNChallengerReputationWithNormalScore(&keeper, helper.Ctx, 1)

		runnerReputation := SetupReputationForRunner(1)
		keeper.SetReputation(helper.Ctx, runnerReputation[0])

		msgServer := k.NewMsgServerImpl(keeper)

		// Create a slice of pointers to types.ClientPublicKey
		reputationPubKeys := make([]*types.ClientPublicKey, len(reputation))
		for i, rep := range reputation {
			// Create a pointer to the current reputation
			repPointer := &types.ClientPublicKey{
				P: rep.PubKey,
				N: 10, // Assuming rep.N is of type int32
			}
			reputationPubKeys[i] = repPointer
		}

		res, err := msgServer.RunnerChallenge(sdk.WrapSDKContext(helper.Ctx), &types.MsgRunnerChallenge{Creator: Challenger_Address, Runner: RunnerPubKey, Clients: reputationPubKeys, Result: "reward"})
		helper.NoError(err)
		helper.Empty(res)

		reputationUpdated, isFound0 := keeper.GetReputation(helper.Ctx, reputationPubKeys[1].P)
		reputationUpdated2, isFound1 := keeper.GetReputation(helper.Ctx, reputationPubKeys[0].P)
		runnerUpdated, isFound2 := keeper.GetReputationsByAddress(helper.Ctx, RunnerAddress)
		challengerUpdated, isFound3 := keeper.GetReputationsByAddress(helper.Ctx, Challenger_Address)
		helper.Require().NotEmpty(isFound0)
		helper.Require().NotEmpty(isFound1)
		helper.Require().NotEmpty(isFound2)
		helper.Require().NotEmpty(isFound3)
		helper.Equal("88.259", reputationUpdated.Score)
		helper.Equal(reputationUpdated.NetEarnings, "1767859udmotus")
		helper.Equal(reputationUpdated2.NetEarnings, "1232140udmotus")
		helper.Equal(runnerUpdated.NetEarnings, "1000000udmotus")
		helper.Equal(challengerUpdated.NetEarnings, "1000000udmotus")

	})
}

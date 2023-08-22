package keeper_test

func (helper *KeeperTestHelper) TestMintRewardCoins() {

	helper.Run("TestMintRewardCoins", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		epochKeeper := helper.App.EpochKeeper

		//helper.FundModuleAcc(types.ModuleName, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000000))))
		//Set up some example epoch data
		epochData := CreateEpochData(&epochKeeper, helper.Ctx)

		epochKeeper.SetEpochData(helper.Ctx, epochData)

		// Call the MintRewardCoins function
		keeper.MintRewardCoins(helper.Ctx, epochData)

		epochData, isFound := epochKeeper.GetEpochData(helper.Ctx)
		helper.Require().NotEmpty(isFound)

		// Check that all LastBlockChallenges fields are set to zero
		helper.Require().Equal(uint64(0), epochData.V2VRXLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.V2VBXLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.V2NBXLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.RunnerLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.ChallengerLastBlockChallenges)
	})

}

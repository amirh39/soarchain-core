package keeper_test

import (
	"soarchain/x/epoch/types"
)

func (helper *KeeperTestHelper) TestMintRewardCoins() {

	helper.Setup()

	helper.Run("TestMintRewardCoins", func() {
		keeper := helper.App.PoaKeeper
		epochKeeper := helper.App.EpochKeeper

		//helper.FundModuleAcc(types.ModuleName, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000000))))
		//Set up some example epoch data
		epochData := types.EpochData{
			TotalEpochs:                   1,
			EpochV2VRX:                    "0udmotus",
			EpochV2VBX:                    "0udmotus",
			EpochV2NBX:                    "0udmotus",
			EpochRunner:                   "0udmotus",
			EpochChallenger:               "0udmotus",
			V2VRXTotalChallenges:          0,
			V2VBXTotalChallenges:          0,
			V2NBXTotalChallenges:          0,
			ChallengerTotalChallenges:     0,
			RunnerTotalChallenges:         0,
			TotalChallengesPrevDay:        0,
			InitialPerChallengeValue:      12.0,
			V2VRXLastBlockChallenges:      100,
			V2VRXPerChallengeValue:        2,
			V2VBXLastBlockChallenges:      150,
			V2VBXPerChallengeValue:        3,
			V2NBXLastBlockChallenges:      200,
			V2NBXPerChallengeValue:        4,
			RunnerLastBlockChallenges:     250,
			RunnerPerChallengeValue:       5,
			ChallengerLastBlockChallenges: 300,
			ChallengerPerChallengeValue:   6,
		}

		epochKeeper.SetEpochData(helper.Ctx, epochData)

		// Call the MintRewardCoins function
		keeper.MintRewardCoins(helper.Ctx, epochData)

		epochData, _ = epochKeeper.GetEpochData(helper.Ctx)

		// Check that all LastBlockChallenges fields are set to zero
		helper.Require().Equal(uint64(0), epochData.V2VRXLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.V2VBXLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.V2NBXLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.RunnerLastBlockChallenges)
		helper.Require().Equal(uint64(0), epochData.ChallengerLastBlockChallenges)
	})

}

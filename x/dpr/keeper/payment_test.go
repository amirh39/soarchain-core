package keeper_test

import (
	"log"

	didtypes "github.com/soar-robotics/soarchain-core/x/did/types"

	epochtypes "github.com/soar-robotics/soarchain-core/x/epoch/types"
)

func (helper *KeeperTestHelper) Test_DistributeRewards() {
	helper.Run("TestDistributeRewards", func() {
		helper.Setup()

		didKeeper := helper.App.DidKeeper
		dprKeeper := helper.App.DprKeeper
		epochKeeper := helper.App.EpochKeeper
		poaKeeper := helper.App.PoaKeeper

		// Set up mock DPRs (play with the inputs of the dpr's)
		dpr1 := SetupDpr(1)
		dpr2 := SetupSecondDpr(1)
		dprKeeper.SetDpr(helper.Ctx, dpr1[0])
		dprKeeper.SetDpr(helper.Ctx, dpr2[0])

		// Create DprInfo for the first DPR
		dprinfo1 := &didtypes.DprInfo{
			Id:      DprId,
			Claimed: "0",
		}

		// Create DprInfo for the second DPR
		dprinfo2 := &didtypes.DprInfo{
			Id:      DprID,
			Claimed: "0",
		}

		// Create new ClientDid with two DprInfos
		newDid := didtypes.ClientDid{
			Id:       Did,
			DprInfos: []*didtypes.DprInfo{dprinfo1, dprinfo2}, // Include both DPR infos
			Address:  ADDRESS,
		}

		didKeeper.SetClientDid(helper.Ctx, newDid)
		did, _ := didKeeper.GetClientDid(helper.Ctx, ADDRESS)
		// Mock the epoch data
		mockedEpochData := epochtypes.EpochData{
			TotalEpochs: 10,
		}
		epochKeeper.SetEpochData(helper.Ctx, mockedEpochData)

		reputation, _ := poaKeeper.GetReputationsByAddress(helper.Ctx, ADDRESS)
		reputation.DprEarnings = "100udmotus"
		poaKeeper.SetReputation(helper.Ctx, reputation)

		// Call distributeRewards
		rewards, err := dprKeeper.DistributeRewards(helper.Ctx, did)

		// Assertions
		helper.Require().NotNil(rewards)
		helper.Require().Nil(err)

		did1, _ := didKeeper.GetClientDid(helper.Ctx, ADDRESS)
		log.Println(did1.DprInfos[0], did1.DprInfos[1])

		log.Println(reputation)

		// Print baby
		log.Println("Distributed Rewards:", rewards)
	})
}

// TODO: we are losing maximum 1 umotus for each claim -> increase precision

package keeper_test

import (
	"log"
	"soarchain/x/dpr/utility"
)

func (helper *KeeperTestHelper) Test_dpr_EndTime() {
	helper.Run("TestdprEndTime", func() {
		helper.Setup()

		blockTime := helper.Ctx.BlockTime()
		log.Println(blockTime)

		endTime, err := utility.CalculateDPREndTime(blockTime, 192)
		log.Println(endTime) //exactly 1 day after BlockTime, the end time of dpr should come
		helper.Require().NotNil(endTime)
		helper.Require().NoError(err)

	})
}

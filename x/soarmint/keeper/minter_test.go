package keeper_test

import (
	"testing"

	"github.com/amirh39/soarchain-core/app/apptesting"

	"github.com/stretchr/testify/suite"
)

type KeeperTestHelper struct {
	apptesting.KeeperTestSuite
}

func (suite *KeeperTestHelper) SetupTest() {
	suite.Setup()
}
func TestKeeperTestHelper(t *testing.T) {
	suite.Run(t, new(KeeperTestHelper))
}

func (helper *KeeperTestHelper) TestMinterGet() {
	helper.Setup()
	keeper := helper.App.MintKeeper
	minter, isFound := keeper.GetMinter(helper.Ctx)
	helper.NotEmpty(minter)
	helper.True(isFound)

}

func (helper *KeeperTestHelper) TestMinterRemove() {
	helper.Setup()
	keeper := helper.App.MintKeeper
	// no need for setting since it is being set in abci.go
	keeper.RemoveMinter(helper.Ctx)
	minter, isFound := keeper.GetMinter(helper.Ctx)
	helper.Empty(minter)
	helper.False(isFound)
}

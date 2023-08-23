package keeper_test

import (
	"soarchain/app/apptesting"
	"testing"

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

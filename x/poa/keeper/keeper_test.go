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

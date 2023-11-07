package keeper_test

import (
	"soarchain/app/apptesting"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

type KeeperTestHelper struct {
	apptesting.KeeperTestSuite
}

func (suite *KeeperTestHelper) SetupTest() {
	suite.Setup()

}
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func RandomAccountAddress() sdk.AccAddress {
	_, _, addr := keyPubAddr()
	return addr
}
func TestKeeperTestHelper(t *testing.T) {
	suite.Run(t, new(KeeperTestHelper))
}

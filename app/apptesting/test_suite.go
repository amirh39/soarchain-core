package apptesting

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/amirh39/soarchain-core/app"
	param "github.com/amirh39/soarchain-core/app/params"

	dprtypes "github.com/amirh39/soarchain-core/x/dpr/types"
	poatypes "github.com/amirh39/soarchain-core/x/poa/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite
	Ctx         sdk.Context
	App         *app.SoarchainApp
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
	MsgServer   dprtypes.MsgServer
}

// Setup sets up basic environment for suite (App, Ctx, and test accounts)
func (s *KeeperTestSuite) Setup() {
	s.App = app.Setup(false)
	s.Ctx = s.App.BaseApp.NewContext(false, tenderminttypes.Header{Height: 1, ChainID: "soarchain", Time: time.Now().UTC()})
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}

	s.TestAccs = CreateRandomAccounts(3)
}

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

func (s *KeeperTestSuite) FundAccount(acc sdk.AccAddress, amounts sdk.Coins) {

	ctx := s.App.NewContext(true, tenderminttypes.Header{Height: s.App.LastBlockHeight()})

	err := simapp.FundAccount(s.App.BankKeeper, ctx, acc, amounts)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) FundModuleAcc(moduleName string, amounts sdk.Coins) {
	err := simapp.FundModuleAccount(s.App.BankKeeper, s.Ctx, moduleName, amounts)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) MintCoins(coins sdk.Coins) {
	err := s.App.BankKeeper.MintCoins(s.Ctx, poatypes.ModuleName, coins)
	s.Require().NoError(err)
}

// CreateTestContextWithMultiStore creates a test context and returns it together with multi store.
func (s *KeeperTestSuite) CreateTestContextWithMultiStore() (sdk.Context, sdk.CommitMultiStore) {
	db := dbm.NewMemDB()
	logger := log.NewNopLogger()

	ms := rootmulti.NewStore(db)

	return sdk.NewContext(ms, tenderminttypes.Header{}, false, logger), ms
}

// CreateTestContext creates a test context.
func (s *KeeperTestSuite) Commit() {
	oldHeight := s.Ctx.BlockHeight()
	oldHeader := s.Ctx.BlockHeader()
	s.App.Commit()
	newHeader := tenderminttypes.Header{Height: oldHeight + 1, ChainID: oldHeader.ChainID, Time: oldHeader.Time.Add(time.Second)}
	s.App.BeginBlock(abci.RequestBeginBlock{Header: newHeader})
	s.Ctx = s.App.GetBaseApp().NewContext(false, newHeader)
}

// SetupValidator sets up a validator and returns the ValAddress.
func (s *KeeperTestSuite) SetupValidatorInvariant(bondStatus stakingtypes.BondStatus) sdk.ValAddress {

	valPubKey := secp256k1.GenPrivKey().PubKey()
	valAddress := sdk.ValAddress(valPubKey.Address())
	bondDenom := param.BondDenom
	selfBond := sdk.NewCoins(sdk.Coin{Amount: sdk.NewInt(100), Denom: bondDenom})

	s.FundAccount(sdk.AccAddress(valAddress), selfBond)

	stakingHandler := staking.NewHandler(s.App.StakingKeeper)
	stakingCoin := sdk.NewCoin(sdk.DefaultBondDenom, selfBond[0].Amount)
	ZeroCommission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	msg, err := stakingtypes.NewMsgCreateValidator(valAddress, valPubKey, stakingCoin, stakingtypes.Description{}, ZeroCommission, sdk.OneInt())
	s.Require().NoError(err)
	res, err := stakingHandler(s.Ctx, msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	val, found := s.App.StakingKeeper.GetValidator(s.Ctx, valAddress)
	s.Require().True(found)

	val = val.UpdateStatus(bondStatus)
	s.App.StakingKeeper.SetValidator(s.Ctx, val)

	consAddress, err := val.GetConsAddr()
	s.Suite.Require().NoError(err)

	signingInfo := slashingtypes.NewValidatorSigningInfo(
		consAddress,
		s.Ctx.BlockHeight(),
		0,
		time.Unix(0, 0),
		false,
		0,
	)
	s.App.SlashingKeeper.SetValidatorSigningInfo(s.Ctx, consAddress, signingInfo)

	return valAddress
}

// SetupMultipleValidators setups "numValidator" validators and returns their address in string
func (s *KeeperTestSuite) SetupMultipleValidators(numValidator int) []string {
	valAddress := []string{}
	for i := 0; i < numValidator; i++ {
		valAddr := s.SetupValidatorInvariant(stakingtypes.Bonded)
		valAddress = append(valAddress, valAddr.String())
	}
	return valAddress
}

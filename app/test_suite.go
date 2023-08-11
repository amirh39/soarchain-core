package app

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	param "soarchain/app/params"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type KeeperTestHelper struct {
	suite.Suite
	Ctx         sdk.Context
	app         *soarchainApp
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
}

func (s *KeeperTestHelper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {

	ctx := s.app.NewContext(true, tmproto.Header{Height: s.app.LastBlockHeight()})

	err := simapp.FundAccount(s.app.BankKeeper, ctx, acc, amounts)
	s.Require().NoError(err)
}

// SetupValidator sets up a validator and returns the ValAddress.
func (s *KeeperTestHelper) SetupValidatorInvarient(bondStatus stakingtypes.BondStatus) sdk.ValAddress {

	valPubKey := secp256k1.GenPrivKey().PubKey()
	valAddress := sdk.ValAddress(valPubKey.Address())
	bondDenom := param.BondDenom
	selfBond := sdk.NewCoins(sdk.Coin{Amount: sdk.NewInt(100), Denom: bondDenom})

	s.FundAcc(sdk.AccAddress(valAddress), selfBond)

	stakingHandler := staking.NewHandler(s.app.StakingKeeper)
	stakingCoin := sdk.NewCoin(sdk.DefaultBondDenom, selfBond[0].Amount)
	ZeroCommission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	msg, err := stakingtypes.NewMsgCreateValidator(valAddress, valPubKey, stakingCoin, stakingtypes.Description{}, ZeroCommission, sdk.OneInt())
	s.Require().NoError(err)
	res, err := stakingHandler(s.Ctx, msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	val, found := s.app.StakingKeeper.GetValidator(s.Ctx, valAddress)
	s.Require().True(found)

	val = val.UpdateStatus(bondStatus)
	s.app.StakingKeeper.SetValidator(s.Ctx, val)

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
	s.app.SlashingKeeper.SetValidatorSigningInfo(s.Ctx, consAddress, signingInfo)

	return valAddress
}

// SetupMultipleValidators setups "numValidator" validators and returns their address in string
func (s *KeeperTestHelper) SetupMultipleValidators(numValidator int) []string {
	valAddress := []string{}
	for i := 0; i < numValidator; i++ {
		valAddr := s.SetupValidatorInvarient(stakingtypes.Bonded)
		valAddress = append(valAddress, valAddr.String())
	}
	return valAddress
}

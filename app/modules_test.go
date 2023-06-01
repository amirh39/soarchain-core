package app

import (
	"encoding/json"
	"testing"

	param "soarchain/app/params"

	"github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/assert"
)

// Test_newGovModule tests that the default genesis state for the gov module
// uses the utmotus denominiation.
func Test_newGovModule(t *testing.T) {
	encCfg := MakeEncodingConfig()

	govModule := newGovModule()
	raw := govModule.DefaultGenesis(encCfg.Marshaler)
	govGenesisState := govtypes.GenesisState{}

	// HACKHACK explicitly ignore the error returned from json.Unmarshal because
	// the error is a failure to unmarshal the string StartingProposalId as a
	// uint which is unrelated to the test here.
	_ = json.Unmarshal(raw, &govGenesisState)

	want := []types.Coin{{
		Denom:  param.BondDenom,
		Amount: types.NewInt(10000000),
	}}

	assert.Equal(t, want, govGenesisState.DepositParams.MinDeposit)
}

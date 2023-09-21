package app

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	param "soarchain/app/params"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// Test_newGovModule tests that the default genesis state for the gov module
// uses the udmotus denominiation.
func Test_newGovModule(t *testing.T) {
	encCfg := MakeEncodingConfig()

	govModule := newGovModule()
	raw := govModule.DefaultGenesis(encCfg.Marshaler)
	govGenesisState := govtypes.GenesisState{}

	// HACKHACK explicitly ignore the error returned from json.Unmarshal because
	// the error is a failure to unmarshal the string StartingProposalId as a
	// uint which is unrelated to the test here.
	_ = json.Unmarshal(raw, &govGenesisState)

	want := sdk.NewCoins(sdk.NewCoin(param.BondDenom, sdk.NewInt(10000000)))

	assert.Equal(t, want, govGenesisState.DepositParams.MinDeposit)
}

func TestOrderEndandBeginBlockers_determinism(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewSoarchainApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, simapp.EmptyAppOptions{}, GetWasmEnabledProposals(), WasmOptions)

	for i := 0; i < 1000; i++ {
		a := OrderEndBlockers(app.mm.ModuleNames())
		b := OrderEndBlockers(app.mm.ModuleNames())
		c := OrderBeginBlockers(app.mm.ModuleNames())
		d := OrderBeginBlockers(app.mm.ModuleNames())
		require.True(t, reflect.DeepEqual(a, b))
		require.True(t, reflect.DeepEqual(c, d))
	}
}

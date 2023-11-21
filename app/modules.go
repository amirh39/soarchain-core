package app

import (
	"encoding/json"

	param "github.com/soar-robotics/soarchain-core/app/params"
	dprmoduletypes "github.com/soar-robotics/soarchain-core/x/dpr/types"
	epochmoduletypes "github.com/soar-robotics/soarchain-core/x/epoch/types"
	poamoduletypes "github.com/soar-robotics/soarchain-core/x/poa/types"
	mint "github.com/soar-robotics/soarchain-core/x/soarmint"
	minttypes "github.com/soar-robotics/soarchain-core/x/soarmint/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcclientclient "github.com/cosmos/ibc-go/v3/modules/core/02-client/client"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	partialorder "github.com/soar-robotics/soarchain-core/utils"
)

type bankModule struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns custom x/bank module genesis state.
func (bankModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	metadata := banktypes.Metadata{
		Description: "The native staking token of the Soarchain-core.",
		Base:        param.BondDenom,
		Name:        param.DisplayDenom,
		Display:     param.DisplayDenom,
		Symbol:      param.DisplayDenom,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    param.BondDenom,
				Exponent: 0,
				Aliases: []string{
					param.BondDenomAlias,
				},
			},
			{
				Denom:    param.DisplayDenom,
				Exponent: param.MotusExponent,
				Aliases:  []string{},
			},
		},
	}

	genState := banktypes.DefaultGenesisState()
	genState.DenomMetadata = append(genState.DenomMetadata, metadata)

	return cdc.MustMarshalJSON(genState)
}

// stakingModule wraps the x/staking module in order to overwrite specific
// ModuleManager APIs.
type stakingModule struct {
	staking.AppModuleBasic
}

// DefaultGenesis returns custom x/staking module genesis state.
func (stakingModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	params := stakingtypes.DefaultParams()
	params.BondDenom = param.BondDenom

	return cdc.MustMarshalJSON(&stakingtypes.GenesisState{
		Params: params,
	})
}

type crisisModule struct {
	crisis.AppModuleBasic
}

// DefaultGenesis returns custom x/crisis module genesis state.
func (crisisModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&crisistypes.GenesisState{
		ConstantFee: sdk.NewCoin(param.BondDenom, sdk.NewInt(1000)),
	})
}

type mintModule struct {
	mint.AppModuleBasic
}

// DefaultGenesis returns custom x/soarmint module genesis state.
func (mintModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := minttypes.DefaultGenesis()
	genState.Params.MintDenom = param.BondDenom

	return cdc.MustMarshalJSON(genState)
}

func newGovModule() govModule {
	return govModule{gov.NewAppModuleBasic(
		append(
			wasmclient.ProposalHandlers,
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
		)...,
	)}
}

type govModule struct {
	gov.AppModuleBasic
}

func (govModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := govtypes.DefaultGenesisState()
	genState.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(param.BondDenom, sdk.NewInt(param.MinDeposit)))

	return cdc.MustMarshalJSON(genState)
}

// orderBeginBlockers returns the order of BeginBlockers, by module name.
func OrderBeginBlockers(allModuleNames []string) []string {
	ord := partialorder.NewPartialOrdering(allModuleNames)
	// Upgrades should be run VERY first
	ord.FirstElements(upgradetypes.ModuleName, minttypes.ModuleName, poamoduletypes.ModuleName, epochmoduletypes.ModuleName, capabilitytypes.ModuleName)
	// Staking ordering
	ord.Sequence(distrtypes.ModuleName, slashingtypes.ModuleName, evidencetypes.ModuleName, stakingtypes.ModuleName)
	// IBChost came after staking
	ord.Sequence(stakingtypes.ModuleName, ibchost.ModuleName)

	ord.LastElements(dprmoduletypes.ModuleName, wasm.ModuleName)
	return ord.TotalOrdering()
}

// OrderEndBlockers returns EndBlockers (crisis, govtypes, staking) with no relative order.
func OrderEndBlockers(allModuleNames []string) []string {
	ord := partialorder.NewPartialOrdering(allModuleNames)

	// Staking must be after gov.
	ord.FirstElements(govtypes.ModuleName)
	ord.LastElements(stakingtypes.ModuleName)

	return ord.TotalOrdering()
}

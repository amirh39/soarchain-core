package app

import (
	"encoding/json"
	param "soarchain/app/params"
	mint "soarchain/x/soarmint"
	minttypes "soarchain/x/soarmint/types"

	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	ibcclientclient "github.com/cosmos/ibc-go/v3/modules/core/02-client/client"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	genState.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(param.BondDenom, sdk.NewInt(10000000)))

	return cdc.MustMarshalJSON(genState)
}
package poa

import (
	"math/rand"

	"soarchain/testutil/sample"
	poasimulation "soarchain/x/poa/simulation"
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = poasimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgGenClient = "op_weight_msg_gen_client"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenClient int = 100

	opWeightMsgGenChallenger = "op_weight_msg_gen_challenger"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenChallenger int = 100

	opWeightMsgChallengeService = "op_weight_msg_challenge_service"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChallengeService int = 100

	opWeightMsgUnregisterClient = "op_weight_msg_unregister_client"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterClient int = 100

	opWeightMsgUnregisterChallenger = "op_weight_msg_unregister_challenger"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterChallenger int = 100

	opWeightMsgGenGuard = "op_weight_msg_gen_guard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenGuard int = 100

	opWeightMsgCreateTotalClients = "op_weight_msg_total_clients"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTotalClients int = 100

	opWeightMsgUpdateTotalClients = "op_weight_msg_total_clients"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTotalClients int = 100

	opWeightMsgDeleteTotalClients = "op_weight_msg_total_clients"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTotalClients int = 100

	opWeightMsgUnregisterRunner = "op_weight_msg_unregister_runner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterRunner int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	poaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&poaGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgGenClient int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenClient, &weightMsgGenClient, nil,
		func(_ *rand.Rand) {
			weightMsgGenClient = defaultWeightMsgGenClient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenClient,
		poasimulation.SimulateMsgGenClient(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChallengeService int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgChallengeService, &weightMsgChallengeService, nil,
		func(_ *rand.Rand) {
			weightMsgChallengeService = defaultWeightMsgChallengeService
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChallengeService,
		poasimulation.SimulateMsgChallengeService(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnregisterClient int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnregisterClient, &weightMsgUnregisterClient, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterClient = defaultWeightMsgUnregisterClient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnregisterClient,
		poasimulation.SimulateMsgUnregisterClient(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnregisterChallenger int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnregisterChallenger, &weightMsgUnregisterChallenger, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterChallenger = defaultWeightMsgUnregisterChallenger
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnregisterChallenger,
		poasimulation.SimulateMsgUnregisterChallenger(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgGenGuard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenGuard, &weightMsgGenGuard, nil,
		func(_ *rand.Rand) {
			weightMsgGenGuard = defaultWeightMsgGenGuard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenGuard,
		poasimulation.SimulateMsgGenGuard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTotalClients int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTotalClients, &weightMsgCreateTotalClients, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTotalClients = defaultWeightMsgCreateTotalClients
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTotalClients,
		poasimulation.SimulateMsgCreateTotalClients(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTotalClients int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTotalClients, &weightMsgUpdateTotalClients, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTotalClients = defaultWeightMsgUpdateTotalClients
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTotalClients,
		poasimulation.SimulateMsgUpdateTotalClients(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTotalClients int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTotalClients, &weightMsgDeleteTotalClients, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTotalClients = defaultWeightMsgDeleteTotalClients
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTotalClients,
		poasimulation.SimulateMsgDeleteTotalClients(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnregisterRunner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnregisterRunner, &weightMsgUnregisterRunner, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterRunner = defaultWeightMsgUnregisterRunner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnregisterRunner,
		poasimulation.SimulateMsgUnregisterRunner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

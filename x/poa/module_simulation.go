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

	opWeightMsgUnregisterRunner = "op_weight_msg_unregister_runner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterRunner int = 100

	opWeightMsgRunnerChallenge = "op_weight_msg_runner_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRunnerChallenge int = 100

	opWeightMsgUnregisterGuard = "op_weight_msg_unregister_guard"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterGuard int = 100

	opWeightMsgSelectRandomChallenger = "op_weight_msg_select_random_challenger"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSelectRandomChallenger int = 100

	opWeightMsgSelectRandomRunner = "op_weight_msg_select_random_runner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSelectRandomRunner int = 100

	opWeightMsgV2VChallenge = "op_weight_msg_v_2_v_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgV2VChallenge int = 100

	opWeightMsgV2NChallenge = "op_weight_msg_v_2_n_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgV2NChallenge int = 100

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

	var weightMsgRunnerChallenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRunnerChallenge, &weightMsgRunnerChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgRunnerChallenge = defaultWeightMsgRunnerChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRunnerChallenge,
		poasimulation.SimulateMsgRunnerChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnregisterGuard int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnregisterGuard, &weightMsgUnregisterGuard, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterGuard = defaultWeightMsgUnregisterGuard
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnregisterGuard,
		poasimulation.SimulateMsgUnregisterGuard(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSelectRandomChallenger int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSelectRandomChallenger, &weightMsgSelectRandomChallenger, nil,
		func(_ *rand.Rand) {
			weightMsgSelectRandomChallenger = defaultWeightMsgSelectRandomChallenger
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSelectRandomChallenger,
		poasimulation.SimulateMsgSelectRandomChallenger(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSelectRandomRunner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSelectRandomRunner, &weightMsgSelectRandomRunner, nil,
		func(_ *rand.Rand) {
			weightMsgSelectRandomRunner = defaultWeightMsgSelectRandomRunner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSelectRandomRunner,
		poasimulation.SimulateMsgSelectRandomRunner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgV2VChallenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgV2VChallenge, &weightMsgV2VChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgV2VChallenge = defaultWeightMsgV2VChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgV2VChallenge,
		poasimulation.SimulateMsgV2VChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgV2NChallenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgV2NChallenge, &weightMsgV2NChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgV2NChallenge = defaultWeightMsgV2NChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgV2NChallenge,
		poasimulation.SimulateMsgV2NChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

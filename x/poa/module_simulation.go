package poa

import (
	"math/rand"

	"github.com/soar-robotics/soarchain-core/testutil/sample"
	poasimulation "github.com/soar-robotics/soarchain-core/x/poa/simulation"
	"github.com/soar-robotics/soarchain-core/x/poa/types"

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
	opWeightMsgGenReputation = "op_weight_msg_gen_reputation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenReputation int = 100

	opWeightMsgChallengeService = "op_weight_msg_challenge_service"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChallengeService int = 100

	opWeightMsgUnregisterReputation = "op_weight_msg_unregister_reputation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnregisterReputation int = 100

	opWeightMsgRunnerChallenge = "op_weight_msg_runner_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRunnerChallenge int = 100

	opWeightMsgSelectRandomChallenger = "op_weight_msg_select_random_challenger"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSelectRandomChallenger int = 100

	opWeightMsgClaimMotusRewards = "op_weight_msg_claim_motus_rewards"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimMotusRewards int = 100

	opWeightMsgClaimRunnerRewards = "op_weight_msg_claim_runner_rewards"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimRunnerRewards int = 100

	opWeightMsgRegisterFactoryKey = "op_weight_msg_register_factory_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterFactoryKey int = 100

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

	var weightMsgClaimMotusRewards int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimMotusRewards, &weightMsgClaimMotusRewards, nil,
		func(_ *rand.Rand) {
			weightMsgClaimMotusRewards = defaultWeightMsgClaimMotusRewards
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimMotusRewards,
		poasimulation.SimulateMsgClaimMotusRewards(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimRunnerRewards int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimRunnerRewards, &weightMsgClaimRunnerRewards, nil,
		func(_ *rand.Rand) {
			weightMsgClaimRunnerRewards = defaultWeightMsgClaimRunnerRewards
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimRunnerRewards,
		poasimulation.SimulateMsgClaimRunnerRewards(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterFactoryKey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterFactoryKey, &weightMsgRegisterFactoryKey, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterFactoryKey = defaultWeightMsgRegisterFactoryKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterFactoryKey,
		poasimulation.SimulateMsgRegisterFactoryKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

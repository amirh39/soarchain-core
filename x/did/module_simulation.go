package did

import (
	"math/rand"

	"github.com/soar-robotics/soarchain-core/testutil/sample"
	didsimulation "github.com/soar-robotics/soarchain-core/x/did/simulation"
	"github.com/soar-robotics/soarchain-core/x/did/types"

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
	_ = didsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgGenClient = "op_weight_msg_gen_client"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenClient int = 100

	opWeightMsgGenRunner = "op_weight_msg_gen_runner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenRunner int = 100

	opWeightMsgGenChallenger = "op_weight_msg_gen_challenger"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenChallenger int = 100

	opWeightMsgUpdateDid = "op_weight_msg_update_did"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDid int = 100

	opWeightMsgDeactivateDid = "op_weight_msg_deactivate_did"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeactivateDid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	didGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&didGenesis)
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
		didsimulation.SimulateMsgGenClient(am.accountKeeper, am.keeper),
	))

	var weightMsgGenRunner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenRunner, &weightMsgGenRunner, nil,
		func(_ *rand.Rand) {
			weightMsgGenRunner = defaultWeightMsgGenRunner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenRunner,
		didsimulation.SimulateMsgGenRunner(am.accountKeeper, am.keeper),
	))

	var weightMsgGenChallenger int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenRunner, &weightMsgGenChallenger, nil,
		func(_ *rand.Rand) {
			weightMsgGenChallenger = defaultWeightMsgGenChallenger
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenChallenger,
		didsimulation.SimulateMsgGenChallenger(am.accountKeeper, am.keeper),
	))

	var weightMsgdeactivateDid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivateDid, &weightMsgdeactivateDid, nil,
		func(_ *rand.Rand) {
			weightMsgdeactivateDid = defaultWeightMsgDeactivateDid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgdeactivateDid,
		didsimulation.SimulateMsgDeactivateDid(am.accountKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

package dpr

import (
	"math/rand"

	"github.com/soar-robotics/soarchain-core/testutil/sample"
	dprsimulation "github.com/soar-robotics/soarchain-core/x/dpr/simulation"
	"github.com/soar-robotics/soarchain-core/x/dpr/types"

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
	_ = dprsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgGenDPR = "op_weight_msg_gen_dpr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenDpr int = 100

	opWeightMsgEnterDPR = "op_weight_msg_enter_dpr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnterDpr int = 100

	opWeightMsgLeaveDPR = "op_weight_msg_leave_dpr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLeaveDpr int = 100

	opWeightMsgActivateDPR = "op_weight_msg_activate_dpr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgActivateDpr int = 100

	opWeightMsgDeactivationDpr = "op_weight_msg_deactivation_dpr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeactivateDpr int = 100

	opWeightMsgUpdateDpr = "op_weight_msg_update_dpr"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDpr int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dprGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dprGenesis)
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

	var weightMsgGenDpr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenDPR, &weightMsgGenDpr, nil,
		func(_ *rand.Rand) {
			weightMsgGenDpr = defaultWeightMsgGenDpr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenDpr,
		dprsimulation.SimulateMsgGenDpr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEnterDpr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnterDPR, &weightMsgEnterDpr, nil,
		func(_ *rand.Rand) {
			weightMsgEnterDpr = defaultWeightMsgEnterDpr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnterDpr,
		dprsimulation.SimulateMsgGenDpr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLeaveDpr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLeaveDPR, &weightMsgLeaveDpr, nil,
		func(_ *rand.Rand) {
			weightMsgLeaveDpr = defaultWeightMsgLeaveDpr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLeaveDpr,
		dprsimulation.SimulateMsgLeaveDpr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgActivateDpr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgActivateDPR, &weightMsgActivateDpr, nil,
		func(_ *rand.Rand) {
			weightMsgActivateDpr = defaultWeightMsgActivateDpr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgActivateDpr,
		dprsimulation.SimulateMsgActivateDpr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeactivateDpr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivationDpr, &weightMsgDeactivateDpr, nil,
		func(_ *rand.Rand) {
			weightMsgDeactivateDpr = defaultWeightMsgDeactivateDpr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeactivateDpr,
		dprsimulation.SimulateMsgActivateDpr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDpr int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDpr, &weightMsgUpdateDpr, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDpr = defaultWeightMsgUpdateDpr
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDpr,
		dprsimulation.SimulateMsgActivateDpr(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

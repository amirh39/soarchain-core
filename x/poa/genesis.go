package poa

import (
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the challenger
	for _, elem := range genState.ChallengerList {
		k.SetChallenger(ctx, elem)
	}
	// Set all the runner
	for _, elem := range genState.RunnerList {
		k.SetRunner(ctx, elem)
	}

	// Set all the vrfData
	for _, elem := range genState.VrfDataList {
		k.SetVrfData(ctx, elem)
	}

	// Set if defined
	k.SetMasterKey(ctx, genState.MasterKey)
	// Set all the factoryKeys
	for _, elem := range genState.FactoryKeysList {
		k.SetFactoryKeys(ctx, elem)
	}

	// Set factoryKeys count
	k.SetFactoryKeysCount(ctx, genState.FactoryKeysCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ChallengerList = k.GetAllChallenger(ctx)
	genesis.RunnerList = k.GetAllRunner(ctx)

	genesis.VrfDataList = k.GetAllVrfData(ctx)
	// Get all masterKey
	masterKey, found := k.GetMasterKey(ctx)
	if found {
		genesis.MasterKey = masterKey
	}
	genesis.FactoryKeysList = k.GetAllFactoryKeys(ctx)
	genesis.FactoryKeysCount = k.GetFactoryKeysCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

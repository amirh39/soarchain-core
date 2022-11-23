package poa

import (
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the client
	for _, elem := range genState.ClientList {
		k.SetClient(ctx, elem)
	}
	// Set all the challenger
	for _, elem := range genState.ChallengerList {
		k.SetChallenger(ctx, elem)
	}
	// Set all the runner
	for _, elem := range genState.RunnerList {
		k.SetRunner(ctx, elem)
	}
	// Set all the guard
	for _, elem := range genState.GuardList {
		k.SetGuard(ctx, elem)
	}
	// Set if defined
	k.SetTotalClients(ctx, genState.TotalClients)
	// Set if defined
	k.SetTotalChallengers(ctx, genState.TotalChallengers)
	// Set if defined
	k.SetTotalRunners(ctx, genState.TotalRunners)

	// Set all the challengerByIndex
	for _, elem := range genState.ChallengerByIndexList {
		k.SetChallengerByIndex(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClientList = k.GetAllClient(ctx)
	genesis.ChallengerList = k.GetAllChallenger(ctx)
	genesis.RunnerList = k.GetAllRunner(ctx)
	genesis.GuardList = k.GetAllGuard(ctx)
	// Get all totalClients
	totalClients, found := k.GetTotalClients(ctx)
	if found {
		genesis.TotalClients = totalClients
	}
	// Get all totalChallengers
	totalChallengers, found := k.GetTotalChallengers(ctx)
	if found {
		genesis.TotalChallengers = totalChallengers
	}
	// Get all totalRunners
	totalRunners, found := k.GetTotalRunners(ctx)
	if found {
		genesis.TotalRunners = totalRunners
	}
	genesis.ChallengerByIndexList = k.GetAllChallengerByIndex(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

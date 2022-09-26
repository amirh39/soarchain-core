package poa

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
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
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClientList = k.GetAllClient(ctx)
	genesis.ChallengerList = k.GetAllChallenger(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

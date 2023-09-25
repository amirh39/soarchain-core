package did

import (
	"soarchain/x/did/keeper"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	// Set all the reputation
	for _, elem := range genState.ReputationList {
		k.SetReputation(ctx, elem)
	}

	for _, element := range genState.DidList {
		did, found := k.GetDidDocument(ctx, element)
		if found {
			k.SetDidDocument(ctx, did.Document.Id, did)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DidList = k.GetAllDid(ctx)
	genesis.ReputationList = k.GetAllReputation(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

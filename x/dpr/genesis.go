package dpr

import (
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// for dpr, doc := range genState.Documents {
	// 	k.SetDidDocument(ctx, dpr, *doc)
	// }

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	//documentsMap := make(map[string]*types.DidDocumentWithSeq)
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	//logger := k.Logger(ctx)

	// for _, dpr := range k.GetAllDid(ctx) {
	// 	key := types.GenesisDidDocumentKey{Did: dpr}.Marshal()
	// 	document, found := k.GetDidDocument(ctx, dpr)
	// 	if !found {
	// 		if logger != nil {
	// 			logger.Error("Exporting Did document state failed through Genesis.", "function", "ExportGenesis")
	// 		}
	// 		return nil
	// 	}
	// 	documentsMap[key] = &document
	// }

	//genesis.Documents = documentsMap
	//genesis.DidList = k.GetAllDid(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

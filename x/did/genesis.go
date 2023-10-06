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

	for _, element := range genState.ClientDidList {
		did, found := k.GetClientDidDocument(ctx, element.Document.Id)
		if found {
			k.SetClientDidDocument(ctx, did.Document.Id, did)
		}
	}

	for _, element := range genState.RunnerDidList {
		did, found := k.GetRunnerDidDocument(ctx, element.Document.Id)
		if found {
			k.SetRunnerDidDocument(ctx, did.Document.Id, did)
		}
	}

	for _, element := range genState.ChallengerDidList {
		did, found := k.GetChallengerDidDocument(ctx, element.Document.Id)
		if found {
			k.SetChallengerDidDocument(ctx, did.Document.Id, did)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClientDidList = k.GetAllClientDid(ctx)
	genesis.RunnerDidList = k.GetAllRunnerDid(ctx)
	genesis.ChallengerDidList = k.GetAllChallengerDid(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

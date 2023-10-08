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
		did, found := k.GetClientDid(ctx, element.Id)
		if found {
			k.SetClientDid(ctx, did)
		}
	}

	for _, element := range genState.RunnerDidList {
		did, found := k.GetRunnerDid(ctx, element.Id)
		if found {
			k.SetRunnerDid(ctx, did)
		}
	}

	for _, element := range genState.ChallengerDidList {
		did, found := k.GetChallengerDid(ctx, element.Id)
		if found {
			k.SetChallengerDid(ctx, did)
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

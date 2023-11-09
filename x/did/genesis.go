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
		k.SetClientDid(ctx, element)

	}

	for _, element := range genState.RunnerDidList {
		k.SetRunnerDid(ctx, element)
	}

	for _, element := range genState.ChallengerDidList {
		k.SetChallengerDid(ctx, element)
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

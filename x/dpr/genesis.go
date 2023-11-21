package dpr

import (
	"github.com/soar-robotics/soarchain-core/x/dpr/keeper"
	"github.com/soar-robotics/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	k.SetParams(ctx, genState.Params)

	for _, elem := range genState.DprList {
		k.SetDpr(ctx, elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all dpr
	genesis.DprList = k.GetAllDpr(ctx)

	return genesis
}

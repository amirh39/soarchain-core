package epoch

import (
	"github.com/soar-robotics/soarchain-core/x/epoch/keeper"
	"github.com/soar-robotics/soarchain-core/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	// Set if defined
	k.SetEpochData(ctx, genState.EpochData)

	k.SetRandomData(ctx, genState.RandomData)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all epochData
	epochData, found := k.GetEpochData(ctx)
	if found {
		genesis.EpochData = epochData
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

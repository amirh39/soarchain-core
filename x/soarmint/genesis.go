package soarmint

import (
	"github.com/amirh39/soarchain-core/x/soarmint/keeper"
	"github.com/amirh39/soarchain-core/x/soarmint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	// Set if defined
	k.SetMinter(ctx, data.Minter)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, data.Params)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all minter
	minter, found := k.GetMinter(ctx)
	if found {
		genesis.Minter = minter
	}
	// this line is used by starport scaffolding # genesis/module/export

	return types.NewGenesisState(minter, genesis.Params)
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	poatypes "soarchain/x/poa/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	// Methods imported from account should be defined here
}

type PoaKeeper interface {
	SetReputation(ctx sdk.Context, reputation poatypes.Reputation)
	GetReputation(ctx sdk.Context, pubkey string) (val poatypes.Reputation, found bool)
	// Methods imported from account should be defined here
}

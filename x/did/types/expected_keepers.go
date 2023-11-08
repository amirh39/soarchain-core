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

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	//GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type PoaKeeper interface {
	InitializeReputation(ctx sdk.Context, reputation poatypes.Reputation, certificate string, runnerStake string, runnerAddress string) error
	InitializeClientReputation(ctx sdk.Context, reputation poatypes.Reputation, certificate string) error
	GetReputation(ctx sdk.Context, pubkey string) (val poatypes.Reputation, found bool)
	RemoveClientReputation(ctx sdk.Context, address string) error
	RemoveRunnerReputation(ctx sdk.Context, creator string) error
	RemoveChallengerReputation(ctx sdk.Context, creator string) error
	GetReputationsByAddress(ctx sdk.Context, address string) (val poatypes.Reputation, found bool)
	// Methods imported from account should be defined here
}

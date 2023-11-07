package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	epochtypes "soarchain/x/epoch/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
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

type EpochKeeper interface {
	GetEpochData(ctx sdk.Context) (epochtypes.EpochData, bool)
	SetEpochData(ctx sdk.Context, epochData epochtypes.EpochData)
	UpdateEpochRewards(ctx sdk.Context, serviceName string, coin sdk.Coin) error
	UpdateEpoch(ctx sdk.Context)
	SetRandomData(ctx sdk.Context, randomData epochtypes.RandomData)
	GetAllRandomNumber(ctx sdk.Context) (list []epochtypes.RandomData)
	RandomNumber(ctx sdk.Context, epochTotal uint64) (string, bool)
}

// type DprKeeper interface {
// 	DistributeRewards(ctx sdk.Context, address string) (sdk.DecCoins, error)
// }

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	didtypes "soarchain/x/did/types"
	epochtypes "soarchain/x/epoch/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
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

type DidKeeper interface {
	SetClientDid(ctx sdk.Context, didDocument didtypes.ClientDid)
	GetClientDid(ctx sdk.Context, address string) (val didtypes.ClientDid, found bool)
	GetEligibleDidByPubkey(ctx sdk.Context, pubkey string) (didDocument didtypes.ClientDid, eligible bool)
	GetEligibleDids(ctx sdk.Context, pins []uint) (found bool)
}

type EpochKeeper interface {
	GetEpochData(ctx sdk.Context) (epochtypes.EpochData, bool)
	SetEpochData(ctx sdk.Context, epochData epochtypes.EpochData)
}

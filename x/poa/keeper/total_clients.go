package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetTotalClients set totalClients in the store
func (k Keeper) SetTotalClients(ctx sdk.Context, totalClients types.TotalClients) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalClientsKey))
	b := k.cdc.MustMarshal(&totalClients)
	store.Set([]byte{0}, b)
}

// GetTotalClients returns totalClients
func (k Keeper) GetTotalClients(ctx sdk.Context) (val types.TotalClients, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalClientsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTotalClients removes totalClients from the store
func (k Keeper) RemoveTotalClients(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalClientsKey))
	store.Delete([]byte{0})
}

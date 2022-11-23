package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetTotalRunners set totalRunners in the store
func (k Keeper) SetTotalRunners(ctx sdk.Context, totalRunners types.TotalRunners) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalRunnersKey))
	b := k.cdc.MustMarshal(&totalRunners)
	store.Set([]byte{0}, b)
}

// GetTotalRunners returns totalRunners
func (k Keeper) GetTotalRunners(ctx sdk.Context) (val types.TotalRunners, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalRunnersKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTotalRunners removes totalRunners from the store
func (k Keeper) RemoveTotalRunners(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalRunnersKey))
	store.Delete([]byte{0})
}

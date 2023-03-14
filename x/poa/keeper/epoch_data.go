package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetEpochData set epochData in the store
func (k Keeper) SetEpochData(ctx sdk.Context, epochData types.EpochData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochDataKey))
	b := k.cdc.MustMarshal(&epochData)
	store.Set([]byte{0}, b)
}

// GetEpochData returns epochData
func (k Keeper) GetEpochData(ctx sdk.Context) (val types.EpochData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochDataKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEpochData removes epochData from the store
func (k Keeper) RemoveEpochData(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochDataKey))
	store.Delete([]byte{0})
}

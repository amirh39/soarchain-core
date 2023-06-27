package keeper

import (
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEpochData set epochData in the store
func (k Keeper) SetEpochData(ctx sdk.Context, epochData types.EpochData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochDataKey))
	b := k.cdc.MustMarshal(&epochData)
	store.Set([]byte{0}, b)
}

// GetEpochData returns epochData
func (k Keeper) GetEpochData(ctx sdk.Context) (types.EpochData, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{0})
	defer iterator.Close()

	if iterator.Valid() {
		var epochData types.EpochData
		k.cdc.MustUnmarshal(iterator.Value(), &epochData)
		return epochData, true
	}

	return types.EpochData{}, false
}

// RemoveEpochData removes epochData from the store
func (k Keeper) RemoveEpochData(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EpochDataKey))
	store.Delete([]byte{0})
}

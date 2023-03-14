package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// GetFactoryKeysCount get the total number of factoryKeys
func (k Keeper) GetFactoryKeysCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FactoryKeysCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFactoryKeysCount set the total number of factoryKeys
func (k Keeper) SetFactoryKeysCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FactoryKeysCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFactoryKeys appends a factoryKeys in the store with a new id and update the count
func (k Keeper) AppendFactoryKeys(
	ctx sdk.Context,
	factoryKeys types.FactoryKeys,
) uint64 {
	// Create the factoryKeys
	count := k.GetFactoryKeysCount(ctx)

	// Set the ID of the appended value
	factoryKeys.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FactoryKeysKey))
	appendedValue := k.cdc.MustMarshal(&factoryKeys)
	store.Set(GetFactoryKeysIDBytes(factoryKeys.Id), appendedValue)

	// Update factoryKeys count
	k.SetFactoryKeysCount(ctx, count+1)

	return count
}

// SetFactoryKeys set a specific factoryKeys in the store
func (k Keeper) SetFactoryKeys(ctx sdk.Context, factoryKeys types.FactoryKeys) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FactoryKeysKey))
	b := k.cdc.MustMarshal(&factoryKeys)
	store.Set(GetFactoryKeysIDBytes(factoryKeys.Id), b)
}

// GetFactoryKeys returns a factoryKeys from its id
func (k Keeper) GetFactoryKeys(ctx sdk.Context, id uint64) (val types.FactoryKeys, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FactoryKeysKey))
	b := store.Get(GetFactoryKeysIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFactoryKeys removes a factoryKeys from the store
func (k Keeper) RemoveFactoryKeys(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FactoryKeysKey))
	store.Delete(GetFactoryKeysIDBytes(id))
}

// GetAllFactoryKeys returns all factoryKeys
func (k Keeper) GetAllFactoryKeys(ctx sdk.Context) (list []types.FactoryKeys) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FactoryKeysKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FactoryKeys
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFactoryKeysIDBytes returns the byte representation of the ID
func GetFactoryKeysIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetFactoryKeysIDFromBytes returns ID in uint64 format from a byte array
func GetFactoryKeysIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

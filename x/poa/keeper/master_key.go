package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soar-robotics/soarchain-core/x/poa/types"
)

// SetMasterKey set masterKey in the store
func (k Keeper) SetMasterKey(ctx sdk.Context, masterKey types.MasterKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MasterKeyKey))
	b := k.cdc.MustMarshal(&masterKey)
	store.Set([]byte{0}, b)
}

// GetMasterKey returns masterKey
func (k Keeper) GetMasterKey(ctx sdk.Context) (val types.MasterKey, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MasterKeyKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMasterKey removes masterKey from the store
func (k Keeper) RemoveMasterKey(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MasterKeyKey))
	store.Delete([]byte{0})
}

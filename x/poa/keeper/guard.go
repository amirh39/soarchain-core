package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetGuard set a specific guard in the store from its index
func (k Keeper) SetGuard(ctx sdk.Context, guard types.Guard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GuardKeyPrefix))
	b := k.cdc.MustMarshal(&guard)
	store.Set(types.GuardKey(
		guard.Index,
	), b)
}

// GetGuard returns a guard from its index
func (k Keeper) GetGuard(
	ctx sdk.Context,
	index string,

) (val types.Guard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GuardKeyPrefix))

	b := store.Get(types.GuardKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGuard removes a guard from the store
func (k Keeper) RemoveGuard(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GuardKeyPrefix))
	store.Delete(types.GuardKey(
		index,
	))
}

// GetAllGuard returns all guard
func (k Keeper) GetAllGuard(ctx sdk.Context) (list []types.Guard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GuardKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Guard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

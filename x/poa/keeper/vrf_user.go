package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetVrfUser set a specific vrfUser in the store from its index
func (k Keeper) SetVrfUser(ctx sdk.Context, vrfUser types.VrfUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfUserKeyPrefix))
	b := k.cdc.MustMarshal(&vrfUser)
	store.Set(types.VrfUserKey(
		vrfUser.Index,
	), b)
}

// GetVrfUser returns a vrfUser from its index
func (k Keeper) GetVrfUser(
	ctx sdk.Context,
	index string,

) (val types.VrfUser, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfUserKeyPrefix))

	b := store.Get(types.VrfUserKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVrfUser removes a vrfUser from the store
func (k Keeper) RemoveVrfUser(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfUserKeyPrefix))
	store.Delete(types.VrfUserKey(
		index,
	))
}

// GetAllVrfUser returns all vrfUser
func (k Keeper) GetAllVrfUser(ctx sdk.Context) (list []types.VrfUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfUserKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VrfUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

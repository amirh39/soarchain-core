package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soar-robotics/soarchain-core/x/poa/types"
)

// SetVrfData set a specific vrfData in the store from its index
func (k Keeper) SetVrfData(ctx sdk.Context, vrfData types.VrfData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfDataKeyPrefix))
	b := k.cdc.MustMarshal(&vrfData)
	store.Set(types.VrfDataKey(
		vrfData.Index,
	), b)
}

// GetVrfData returns a vrfData from its index
func (k Keeper) GetVrfData(
	ctx sdk.Context,
	index string,

) (val types.VrfData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfDataKeyPrefix))

	b := store.Get(types.VrfDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVrfData removes a vrfData from the store
func (k Keeper) RemoveVrfData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfDataKeyPrefix))
	store.Delete(types.VrfDataKey(
		index,
	))
}

// GetAllVrfData returns all vrfData
func (k Keeper) GetAllVrfData(ctx sdk.Context) (list []types.VrfData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VrfDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VrfData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

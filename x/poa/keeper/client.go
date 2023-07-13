package keeper

import (
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetClient set a specific client in the store from its index
func (k Keeper) SetClient(ctx sdk.Context, client types.Client) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientKeyPrefix))
	b := k.cdc.MustMarshal(&client)
	store.Set(types.ClientKey(
		client.Index,
	), b)
}

// GetClient returns a client from its index
func (k Keeper) GetClient(
	ctx sdk.Context,
	index string,

) (val types.Client, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientKeyPrefix))

	b := store.Get(types.ClientKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClient removes a client from the store
func (k Keeper) RemoveClient(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientKeyPrefix))
	store.Delete(types.ClientKey(
		index,
	))
}

// GetAllClient returns all client
func (k Keeper) GetAllClient(ctx sdk.Context) (list []types.Client) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Client
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

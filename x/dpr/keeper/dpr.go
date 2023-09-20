package keeper

import (
	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Set Dpr object in the store
func (k Keeper) SetDpr(ctx sdk.Context, dpr types.Dpr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))
	b := k.cdc.MustMarshal(&dpr)
	store.Set(types.DprKey(
		dpr.Id,
	), b)
}

func (k Keeper) GetDpr(ctx sdk.Context, id string) (val types.Dpr, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))

	b := store.Get(types.DprKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllDpr(ctx sdk.Context) (list []types.Dpr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Dpr
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetAllActiveDpr(ctx sdk.Context) (list []types.Dpr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Dpr
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.IsActive {
			list = append(list, val)
		}
	}
	return
}

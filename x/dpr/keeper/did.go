package keeper

import (
	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Set DprDocument in the store
func (k Keeper) SetDpr(ctx sdk.Context, id string, doc types.QueryGetDprRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))
	key := []byte(id)
	bz := k.cdc.MustMarshalLengthPrefixed(&doc)
	store.Set(key, bz)
}

func (k Keeper) GetDprRequest(ctx sdk.Context, id string) (val types.QueryGetDprResponse, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))
	key := []byte(id)
	bz := store.Get(key)
	if bz == nil {
		return val, false
	}
	var doc types.QueryGetDprResponse
	k.cdc.MustUnmarshalLengthPrefixed(bz, &doc)
	return doc, true
}

func (k Keeper) GetAllDpr(ctx sdk.Context) []string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DprKeyPrefix))
	dprs := make([]string, 0)

	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		dpr := string(iter.Key())
		dprs = append(dprs, dpr)
	}
	return dprs
}

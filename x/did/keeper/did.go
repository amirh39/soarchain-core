package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Set DidDocument in the store
func (k Keeper) SetDidDocument(ctx sdk.Context, id string, doc types.DidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	key := []byte(id)
	bz := k.cdc.MustMarshalLengthPrefixed(&doc)
	store.Set(key, bz)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, id string) (val types.DidDocumentWithSeq, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	key := []byte(id)
	bz := store.Get(key)
	if bz == nil {
		return val, false
	}
	var doc types.DidDocumentWithSeq
	k.cdc.MustUnmarshalLengthPrefixed(bz, &doc)
	return doc, true
}

func (k Keeper) GetAllDid(ctx sdk.Context) []string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	dids := make([]string, 0)

	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		did := string(iter.Key())
		dids = append(dids, did)
	}
	return dids
}

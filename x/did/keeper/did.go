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

func (k Keeper) GetDidDocumentWithSequence(ctx sdk.Context, id string) (val types.DidDocumentWithSeq, found bool) {
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

func (k Keeper) GetDidDocumentByPubkey(ctx sdk.Context, pubkey string) (didDocument types.DidDocumentWithSeq, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocumentWithSeq
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &val)
		if val.Document.ClientPublicKey == pubkey {
			return val, true
		}
	}
	return types.DidDocumentWithSeq{}, false
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

func (k Keeper) FindEligibleDid(ctx sdk.Context, pins []uint) (found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocumentWithSeq
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &val)
		switch len(pins) {
		case 1:
			if val.Document.PidSupportedOneToTwnety {
				return true
			}
			if val.Document.PidSupportedTwentyOneToForthy {
				return true
			}
			if val.Document.PidSupportedForthyOneToSixty {
				return true
			}
		case 2:
			if val.Document.PidSupportedOneToTwnety && val.Document.PidSupportedTwentyOneToForthy {
				return true
			}
			if val.Document.PidSupportedTwentyOneToForthy && val.Document.PidSupportedForthyOneToSixty {
				return true
			}
			if val.Document.PidSupportedOneToTwnety && val.Document.PidSupportedForthyOneToSixty {
				return true
			}
		case 3:
			if val.Document.PidSupportedOneToTwnety && val.Document.PidSupportedTwentyOneToForthy && val.Document.PidSupportedForthyOneToSixty {
				return true
			}
		}
	}
	return false
}

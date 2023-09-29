package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetDidDocument(ctx sdk.Context, id string, didDocument types.DidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidKey(
		id,
	), b)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, id string) (val types.DidDocumentWithSeq, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))

	b := store.Get(types.DidKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetEligibleDidByPubkey(ctx sdk.Context, pubkey string) (didDocument types.DidDocumentWithSeq, eligible bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocumentWithSeq
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Document.Index == pubkey {
			return val, true
		}
	}
	return types.DidDocumentWithSeq{}, false
}

func (k Keeper) GetAllDid(ctx sdk.Context) (list []types.DidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocumentWithSeq
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetEligibleDids(ctx sdk.Context, pins []uint) (found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocumentWithSeq
		k.cdc.MustUnmarshal(iterator.Value(), &val)
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

func (k Keeper) ValidateDid(ctx sdk.Context, id string, address string, pubkey string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocumentWithSeq
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Document.Id == id || val.Document.Address == address || val.Document.Index == pubkey {
			return false
		}
	}
	return true
}

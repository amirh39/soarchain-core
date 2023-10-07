package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetClientDidDocument(ctx sdk.Context, id string, didDocument types.ClientDidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidKey(
		id,
	), b)
}

func (k Keeper) GetClientDidDocument(ctx sdk.Context, id string) (val types.ClientDidDocumentWithSeq, found bool) {
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

func (k Keeper) GetEligibleDidByPubkey(ctx sdk.Context, pubkey string) (didDocument types.ClientDidDocumentWithSeq, eligible bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientDidDocumentWithSeq
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Document.PubKey == pubkey {
			return val, true
		}
	}
	return types.ClientDidDocumentWithSeq{}, false
}

func (k Keeper) GetAllClientDid(ctx sdk.Context) (list []types.ClientDidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientDidDocumentWithSeq
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
		var val types.ClientDidDocumentWithSeq
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

func (k Keeper) RemoveClientDid(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	store.Delete(types.DidKey(
		id,
	))
}

package keeper

import (
	"github.com/amirh39/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetClientDid(ctx sdk.Context, didDocument types.ClientDid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.DidKey(
		didDocument.Address,
	), b)
}

func (k Keeper) GetClientDid(ctx sdk.Context, Address string) (val types.ClientDid, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	b := store.Get(types.DidKey(
		Address,
	))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetClientDidId(ctx sdk.Context, id string) (clientDid types.ClientDid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Id == id {
			return val, true
		}
	}

	return types.ClientDid{}, false
}

func (k Keeper) GetEligibleDidByPubkey(ctx sdk.Context, pubkey string) (didDocument types.ClientDid, eligible bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.PubKey == pubkey {
			return val, true
		}
	}
	return types.ClientDid{}, false
}

func (k Keeper) GetAllClientDid(ctx sdk.Context) (list []types.ClientDid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveClientDid(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	store.Delete(types.DidKey(
		address,
	))
}

func (k Keeper) GetClientDidUsingPubKey(ctx sdk.Context, pubKey string) (client types.ClientDid, found bool) {
	runners := k.GetAllClientDid(ctx)

	for _, c := range runners {
		if c.PubKey == pubKey {
			return c, true
		}
	}

	return client, false
}

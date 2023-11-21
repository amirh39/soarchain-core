package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetChallengerDid(ctx sdk.Context, didDocument types.ChallengerDid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.ChallengerDidKey(
		didDocument.Address,
	), b)
}

func (k Keeper) GetChallengerDid(ctx sdk.Context, Address string) (val types.ChallengerDid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))

	b := store.Get(types.ChallengerDidKey(
		Address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetChallengerDidId(ctx sdk.Context, id string) (clientDid types.ChallengerDid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChallengerDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Id == id {
			return val, true
		}
	}

	return types.ChallengerDid{}, false
}

func (k Keeper) GetAllChallengerDid(ctx sdk.Context) (list []types.ChallengerDid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChallengerDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveChallengerDid(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	store.Delete(types.ChallengerDidKey(
		address,
	))
}

func (k Keeper) GetChallengerDidUsingPubKey(ctx sdk.Context, pubKey string) (runner types.ChallengerDid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChallengerDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.PubKey == pubKey {
			return val, true
		}
	}
	return types.ChallengerDid{}, false
}

package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRunnerDid(ctx sdk.Context, didDocument types.RunnerDid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.RunnerDidKey(
		didDocument.Address,
	), b)
}

func (k Keeper) GetRunnerDid(ctx sdk.Context, Address string) (val types.RunnerDid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))

	b := store.Get(types.RunnerDidKey(
		Address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetRunnerDidId(ctx sdk.Context, id string) (clientDid types.RunnerDid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RunnerDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Id == id {
			return val, true
		}
	}

	return types.RunnerDid{}, false
}

func (k Keeper) GetAllRunnerDid(ctx sdk.Context) (list []types.RunnerDid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RunnerDid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveRunnerDid(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))
	store.Delete(types.RunnerDidKey(
		id,
	))
}

func (k Keeper) GetRunnerDidUsingPubKey(ctx sdk.Context, pubKey string) (runner types.RunnerDid, found bool) {
	runners := k.GetAllRunnerDid(ctx)

	for _, c := range runners {
		if c.PubKey == pubKey {
			return c, true
		}
	}

	return runner, false
}

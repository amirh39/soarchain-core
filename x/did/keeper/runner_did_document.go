package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRunnerDidDocument(ctx sdk.Context, id string, didDocument types.RunnerDidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.RunnerDidKey(
		id,
	), b)
}

func (k Keeper) GetRunnerDidDocument(ctx sdk.Context, id string) (val types.RunnerDidDocumentWithSeq, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))

	b := store.Get(types.RunnerDidKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllRunnerDid(ctx sdk.Context) (list []types.RunnerDidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RunnerDidDocumentWithSeq
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

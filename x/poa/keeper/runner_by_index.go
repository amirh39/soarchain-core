package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetRunnerByIndex set a specific runnerByIndex in the store from its index
func (k Keeper) SetRunnerByIndex(ctx sdk.Context, runnerByIndex types.RunnerByIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerByIndexKeyPrefix))
	b := k.cdc.MustMarshal(&runnerByIndex)
	store.Set(types.RunnerByIndexKey(
		runnerByIndex.Index,
	), b)
}

// GetRunnerByIndex returns a runnerByIndex from its index
func (k Keeper) GetRunnerByIndex(
	ctx sdk.Context,
	index string,

) (val types.RunnerByIndex, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerByIndexKeyPrefix))

	b := store.Get(types.RunnerByIndexKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRunnerByIndex removes a runnerByIndex from the store
func (k Keeper) RemoveRunnerByIndex(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerByIndexKeyPrefix))
	store.Delete(types.RunnerByIndexKey(
		index,
	))
}

// GetAllRunnerByIndex returns all runnerByIndex
func (k Keeper) GetAllRunnerByIndex(ctx sdk.Context) (list []types.RunnerByIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerByIndexKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RunnerByIndex
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

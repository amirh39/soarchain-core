package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetChallengerByIndex set a specific challengerByIndex in the store from its index
func (k Keeper) SetChallengerByIndex(ctx sdk.Context, challengerByIndex types.ChallengerByIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerByIndexKeyPrefix))
	b := k.cdc.MustMarshal(&challengerByIndex)
	store.Set(types.ChallengerByIndexKey(
		challengerByIndex.Index,
	), b)
}

// GetChallengerByIndex returns a challengerByIndex from its index
func (k Keeper) GetChallengerByIndex(
	ctx sdk.Context,
	index string,

) (val types.ChallengerByIndex, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerByIndexKeyPrefix))

	b := store.Get(types.ChallengerByIndexKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChallengerByIndex removes a challengerByIndex from the store
func (k Keeper) RemoveChallengerByIndex(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerByIndexKeyPrefix))
	store.Delete(types.ChallengerByIndexKey(
		index,
	))
}

// GetAllChallengerByIndex returns all challengerByIndex
func (k Keeper) GetAllChallengerByIndex(ctx sdk.Context) (list []types.ChallengerByIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerByIndexKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChallengerByIndex
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

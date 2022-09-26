package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetChallenger set a specific challenger in the store from its index
func (k Keeper) SetChallenger(ctx sdk.Context, challenger types.Challenger) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))
	b := k.cdc.MustMarshal(&challenger)
	store.Set(types.ChallengerKey(
		challenger.Index,
	), b)
}

// GetChallenger returns a challenger from its index
func (k Keeper) GetChallenger(
	ctx sdk.Context,
	index string,

) (val types.Challenger, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))

	b := store.Get(types.ChallengerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChallenger removes a challenger from the store
func (k Keeper) RemoveChallenger(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))
	store.Delete(types.ChallengerKey(
		index,
	))
}

// GetAllChallenger returns all challenger
func (k Keeper) GetAllChallenger(ctx sdk.Context) (list []types.Challenger) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Challenger
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

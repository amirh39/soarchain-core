package keeper

import (
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetChallenger set a specific challenger in the store from its index
func (k Keeper) SetChallenger(ctx sdk.Context, challenger types.Challenger) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))
	b := k.cdc.MustMarshal(&challenger)
	store.Set(types.ChallengerKey(
		challenger.Address,
	), b)
}

// GetChallenger returns a challenger from its index
func (k Keeper) GetChallenger(
	ctx sdk.Context,
	Address string,

) (val types.Challenger, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))

	b := store.Get(types.ChallengerKey(
		Address,
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
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))
	store.Delete(types.ChallengerKey(
		address,
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

func (k Keeper) GetChallengerUsingPubKey(ctx sdk.Context, pubKey string) (challenger types.Challenger, found bool) {
	challengers := k.GetAllChallenger(ctx)

	for _, c := range challengers {
		if c.PubKey == pubKey {
			return c, true
		}
	}

	return challenger, false
}

// GetChallenger returns a challenger from its index and type = "v2n"
func (k Keeper) GetChallengerByType(
	ctx sdk.Context,
	Address string,
	Type string,

) (val types.Challenger, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerKeyPrefix))

	b := store.Get(types.ChallengerKey(
		Address,
	))

	k.cdc.MustUnmarshal(b, &val)
	if b == nil || val.Type != Type {
		return val, false
	}

	return val, true
}

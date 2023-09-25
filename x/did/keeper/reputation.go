package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetReputation(ctx sdk.Context, reputation types.Reputation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	b := k.cdc.MustMarshal(&reputation)
	store.Set(types.ReputationKey(
		reputation.Index,
	), b)
}

func (k Keeper) GetReputation(
	ctx sdk.Context,
	pubkey string,

) (val types.Reputation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))

	b := store.Get(types.ReputationKey(
		pubkey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllReputation(ctx sdk.Context) (list []types.Reputation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetReputationByClientAddress(
	ctx sdk.Context,
	address string,

) (val types.Reputation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Address == address {
			return val, true
		}
	}
	return types.Reputation{}, false
}

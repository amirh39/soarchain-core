package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/epoch/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRandomData set randomData in the store
func (k Keeper) SetRandomData(ctx sdk.Context, randomData types.RandomData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandKeyPrefix))
	id := randomData.EpochNumber
	b := k.cdc.MustMarshal(&randomData)
	store.Set(types.RandKey(
		id,
	), b)
}

// GetRandomData returns Random Data
func (k Keeper) GetRandomData(ctx sdk.Context, eapochNumber string) (val types.RandomData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandKeyPrefix))

	b := store.Get(types.RandKey(
		eapochNumber,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllRandomNumber(ctx sdk.Context) (list []types.RandomData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RandKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RandomData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

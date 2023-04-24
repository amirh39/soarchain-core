package keeper

import (
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRunner set a specific runner in the store from its index
func (k Keeper) SetRunner(ctx sdk.Context, runner types.Runner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerKeyPrefix))
	b := k.cdc.MustMarshal(&runner)
	store.Set(types.RunnerKey(
		runner.Address,
	), b)
}

// GetRunner returns a runner from its index
func (k Keeper) GetRunner(
	ctx sdk.Context,
	index string,

) (val types.Runner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerKeyPrefix))

	b := store.Get(types.RunnerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRunner removes a runner from the store
func (k Keeper) RemoveRunner(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerKeyPrefix))
	store.Delete(types.RunnerKey(
		index,
	))
}

// GetAllRunner returns all runner
func (k Keeper) GetAllRunner(ctx sdk.Context) (list []types.Runner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RunnerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Runner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetRunnerUsingPubKey(ctx sdk.Context, pubKey string) (runner types.Runner, found bool) {
	runners := k.GetAllRunner(ctx)

	for _, c := range runners {
		if c.PubKey == pubKey {
			return c, true
		}
	}

	return runner, false
}

package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/soarmint/types"
)

// SetMinter set minter in the store
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
	b := k.cdc.MustMarshal(&minter)
	store.Set([]byte{0}, b)
}

// GetMinter returns minter
func (k Keeper) GetMinter(ctx sdk.Context) (val types.Minter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMinter removes minter from the store
func (k Keeper) RemoveMinter(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
	store.Delete([]byte{0})
}

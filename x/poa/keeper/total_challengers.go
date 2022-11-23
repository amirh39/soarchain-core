package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"soarchain/x/poa/types"
)

// SetTotalChallengers set totalChallengers in the store
func (k Keeper) SetTotalChallengers(ctx sdk.Context, totalChallengers types.TotalChallengers) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalChallengersKey))
	b := k.cdc.MustMarshal(&totalChallengers)
	store.Set([]byte{0}, b)
}

// GetTotalChallengers returns totalChallengers
func (k Keeper) GetTotalChallengers(ctx sdk.Context) (val types.TotalChallengers, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalChallengersKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTotalChallengers removes totalChallengers from the store
func (k Keeper) RemoveTotalChallengers(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TotalChallengersKey))
	store.Delete([]byte{0})
}

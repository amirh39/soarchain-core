package keeper

import (
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMotusWallet set a specific motusWallet in the store from its index
func (k Keeper) SetMotusWallet(ctx sdk.Context, motusWallet types.MotusWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MotusWalletKeyPrefix))
	b := k.cdc.MustMarshal(&motusWallet)
	store.Set(types.MotusWalletKey(
		motusWallet.Address,
	), b)
}

// GetMotusWallet returns a motusWallet from its index
func (k Keeper) GetMotusWallet(
	ctx sdk.Context,
	index string,

) (val types.MotusWallet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MotusWalletKeyPrefix))

	b := store.Get(types.MotusWalletKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMotusWallet removes a motusWallet from the store
func (k Keeper) RemoveMotusWallet(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MotusWalletKeyPrefix))
	store.Delete(types.MotusWalletKey(
		index,
	))
}

// GetAllMotusWallet returns all motusWallet
func (k Keeper) GetAllMotusWallet(ctx sdk.Context) (list []types.MotusWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MotusWalletKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MotusWallet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

package keeper

import (
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetChallengerDidDocument(ctx sdk.Context, id string, didDocument types.ChallengerDidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	b := k.cdc.MustMarshal(&didDocument)
	store.Set(types.ChallengerDidKey(
		id,
	), b)
}

func (k Keeper) GetChallengerDidDocument(ctx sdk.Context, id string) (val types.ChallengerDidDocumentWithSeq, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))

	b := store.Get(types.ChallengerDidKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllChallengerDid(ctx sdk.Context) (list []types.ChallengerDidDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChallengerDidDocumentWithSeq
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveChallengerDid(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengerDidKeyPrefix))
	store.Delete(types.ChallengerDidKey(
		id,
	))
}

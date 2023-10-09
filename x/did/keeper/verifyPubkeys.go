package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsUniquePubKey(k msgServer, ctx sdk.Context, pubkey string) (isUniquePubkey bool) {
	_, isFoundAsClient := k.GetClientDidUsingPubKey(ctx, pubkey)
	_, isFoundAsRunner := k.GetRunnerDidUsingPubKey(ctx, pubkey)
	_, isFoundAsChallenger := k.GetChallengerDidUsingPubKey(ctx, pubkey)
	if isFoundAsClient || isFoundAsChallenger || isFoundAsRunner {
		return true
	}
	return false
}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsUniquePubKey(k msgServer, ctx sdk.Context, address string, pubkey string) (isUniquePubkey bool) {
	_, isFoundAsRunner := k.GetRunnerUsingPubKey(ctx, pubkey)
	_, isFoundAsChallenger := k.GetChallengerUsingPubKey(ctx, pubkey)
	if isFoundAsChallenger || isFoundAsRunner {
		return true
	}
	return false
}

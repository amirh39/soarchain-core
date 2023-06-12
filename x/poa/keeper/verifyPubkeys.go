package keeper

import (

	// "encoding/pem"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsUniquePubKey(k msgServer, ctx sdk.Context, address string, pubkey string) (isUnique bool) {

	_, isFoundWallet := k.GetMotusWallet(ctx, address)
	_, isFoundAsClient := k.GetClient(ctx, pubkey)
	_, isFoundAsRunner := k.GetRunnerUsingPubKey(ctx, pubkey)
	_, isFoundAsChallenger := k.GetChallengerUsingPubKey(ctx, pubkey)
	if isFoundWallet || isFoundAsChallenger || isFoundAsRunner || isFoundAsClient {
		return false
	}
	return true
}

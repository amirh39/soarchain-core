package keeper

import (

	// "encoding/pem"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsUniqueAddress(k msgServer, ctx sdk.Context, address string) (isUniqueAddress bool) {
	_, isFoundAsClient := k.GetClientDid(ctx, address)
	_, isFoundAsRunner := k.GetRunnerDid(ctx, address)
	_, isFoundAsChallenger := k.GetChallengerDid(ctx, address)
	if isFoundAsClient || isFoundAsChallenger || isFoundAsRunner {
		return true
	}
	return false
}

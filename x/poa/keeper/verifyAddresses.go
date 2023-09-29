package keeper

import (

	// "encoding/pem"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsUniqueAddress(k msgServer, ctx sdk.Context, address string) (isUniqueAddress bool) {

	_, isFoundAsRunner := k.GetRunner(ctx, address)
	_, isFoundAsChallenger := k.GetChallenger(ctx, address)
	if isFoundAsChallenger || isFoundAsRunner {
		return true
	}
	return false
}

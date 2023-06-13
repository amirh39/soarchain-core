package keeper

import (

	// "encoding/pem"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IsUniqueAddress(k msgServer, ctx sdk.Context, address string) (isUniqueAddress bool) {

	isFoundAsClient := getClientByAddress(k, ctx, address)
	_, isFoundAsRunner := k.GetRunner(ctx, address)
	_, isFoundAsChallenger := k.GetChallenger(ctx, address)
	if isFoundAsChallenger || isFoundAsRunner || isFoundAsClient {
		return true
	}
	return false
}

func getClientByAddress(k msgServer, ctx sdk.Context, address string) (isFoundAsClient bool) {

	clients := k.GetAllClient(ctx)

	for _, client := range clients {
		if address == client.Address {
			return true
		}
	}
	return false
}

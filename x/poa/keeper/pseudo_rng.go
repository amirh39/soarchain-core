package keeper

import (
	"context"

	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GenerateRandomNumber(goCtx context.Context) int {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allChallengers := k.GetAllChallenger(ctx)

	rand.Seed(ctx.BlockTime().UnixNano())
	min := 0
	max := int(len(allChallengers) - 1)
	n := rand.Intn(max-min+1) + min

	return n
}

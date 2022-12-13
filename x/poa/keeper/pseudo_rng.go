package keeper

import (
	"context"

	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GenerateRandomNumber(goCtx context.Context) int {
	ctx := sdk.UnwrapSDKContext(goCtx)
	totalChallengers, _ := k.GetTotalChallengers(ctx)
	rand.Seed(ctx.BlockTime().UnixNano())
	min := 1
	max := int(totalChallengers.Count)
	n := rand.Intn(max-min+1) + min

	return n
}

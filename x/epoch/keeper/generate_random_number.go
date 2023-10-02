package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"math/rand"
)

func (k Keeper) RandomNumber(ctx sdk.Context, epochTotal uint64) (string, bool) {

	source := rand.NewSource(int64(epochTotal))
	rand_source := rand.New(source)
	randomNumber := rand_source.Int()
	fmt.Println(randomNumber)

	return strconv.Itoa(randomNumber), true
}

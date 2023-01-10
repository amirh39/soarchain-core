package utility

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TeamReserveEmissionPerBlock(ctx sdk.Context) float64 {
	blocksPerMinute := 1.0
	currentBlockNumber := int(ctx.BlockHeight())

	yearsSinceStart := int(math.Floor(float64(currentBlockNumber) / (blocksPerMinute * minutesPerYear)))

	var tokensPerBlock float64
	if yearsSinceStart >= 0 && yearsSinceStart <= 5 {
		tokensPerYear := 57000000.0
		tokensPerBlock = tokensPerYear / (blocksPerMinute * minutesPerYear)

	} else {
		tokensPerBlock = 0
	}

	return tokensPerBlock
}

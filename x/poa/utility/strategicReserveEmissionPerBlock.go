package utility

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func StrategicReserveEmissionPerBlock(ctx sdk.Context) float64 {
	blocksPerMinute := 1.0
	currentBlockNumber := int(ctx.BlockHeight())

	yearsSinceStart := int(math.Floor(float64(currentBlockNumber) / (blocksPerMinute * minutesPerYear)))

	var tokensPerBlock float64
	if yearsSinceStart >= 1 && yearsSinceStart <= 6 {
		tokensPerYear := 46455000000000.0
		tokensPerBlock = tokensPerYear / (blocksPerMinute * minutesPerYear)

	} else {
		tokensPerBlock = 0
	}

	return tokensPerBlock
}

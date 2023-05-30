package utility

import (
	"math"
	"soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func StrategicReserveEmissionPerBlock(ctx sdk.Context) float64 {
	blocksPerMinute := 1.0
	currentBlockNumber := int(ctx.BlockHeight())

	yearsSinceStart := int(math.Floor(float64(currentBlockNumber) / (blocksPerMinute * minutesPerYear)))

	var tokensPerBlock float64
	if yearsSinceStart >= 1 && yearsSinceStart <= 6 {
		tokensPerYear := utilConstants.StrategicReserveEmissionTokensPerYear
		tokensPerBlock = tokensPerYear / (blocksPerMinute * minutesPerYear)

	} else {
		tokensPerBlock = 0
	}

	return tokensPerBlock
}

package utility

import (
	"math"
	"soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TeamReserveEmissionPerBlock(ctx sdk.Context) float64 {
	blocksPerMinute := utilConstants.BlocksPerMinute
	currentBlockNumber := int(ctx.BlockHeight())

	yearsSinceStart := int(math.Floor(float64(currentBlockNumber) / (blocksPerMinute * minutesPerYear)))

	var tokensPerBlock float64
	if yearsSinceStart >= 0 && yearsSinceStart <= 5 {
		tokensPerYear := utilConstants.TeamReserveEmissionTokensPerYear
		tokensPerBlock = tokensPerYear / (blocksPerMinute * minutesPerYear)

	} else {
		tokensPerBlock = 0
	}

	return tokensPerBlock
}

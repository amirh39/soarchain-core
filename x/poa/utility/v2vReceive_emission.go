package utility

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	minutesPerYear = 525600
)

func V2VReceiveRewardEmissionPerBlock(ctx sdk.Context) float64 {

	// ctx := sdk.UnwrapSDKContext(goCtx)

	blocksPerMinute := 1.0
	currentBlockNumber := int(ctx.BlockHeight())

	initialTokensPerYear := 37366809.21 // initial value

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / (blocksPerMinute * minutesPerYear)

	// Calculate the number of years that have passed since the start of the token issuance
	yearsSinceStart := int(math.Floor(float64(currentBlockNumber) / (blocksPerMinute * minutesPerYear)))

	// Calculate the number of times the token issuance rate has been halved
	halvings := int(math.Floor(float64(yearsSinceStart) / 3))

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < halvings; i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / (blocksPerMinute * minutesPerYear)
	}

	return tokensPerBlock
}

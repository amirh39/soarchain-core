package utility

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ChallengerRewardEmissionPerBlock(ctx sdk.Context) (float64, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerMinute := 1.0
	currentBlockNumber := int(ctx.BlockHeight())

	initialTokensPerYear := 21850083.35 // v2n receiver (runner) initial annual emission

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

	return tokensPerBlock, nil
}

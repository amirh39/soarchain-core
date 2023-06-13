package utility

import (
	"math"
	"soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ChallengerRewardEmissionPerBlock(ctx sdk.Context) (float64, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerMinute := utilConstants.BlocksPerMinute
	currentBlockNumber := int(ctx.BlockHeight())

	initialTokensPerYear := utilConstants.ChallengerInitialTokenPerYear // v2n receiver (runner) initial annual emission

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

func ChallengerRewardEmissionPerEpoch(ctx sdk.Context) (float64, error) {
	blocksPerMinute := utilConstants.BlocksPerMinute
	currentBlockNumber := int(ctx.BlockHeight())
	blocksPerEpoch := utilConstants.BlocksPerEpoch
	initialTokensPerYear := utilConstants.ChallengerInitialTokenPerYear

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / (blocksPerMinute * minutesPerYear)

	// Calculate the number of years that have passed since the start of the token issuance
	yearsSinceStart := (currentBlockNumber) / int(blocksPerMinute*minutesPerYear)

	// Calculate the number of times the token issuance rate has been halved
	halvings := (yearsSinceStart) / 3

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < halvings; i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / (blocksPerMinute * minutesPerYear)
	}

	return float64(tokensPerBlock) * blocksPerEpoch, nil

}

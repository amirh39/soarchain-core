package utility

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	minutesPerYear = 525600
)

func V2VRewardEmissionPerBlock(ctx sdk.Context, clientCommunicationMode string) (float64, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerMinute := 1.0
	currentBlockNumber := int(ctx.BlockHeight())

	var initialTokensPerYear float64
	switch clientCommunicationMode {
	case "v2v-rx":
		initialTokensPerYear = 37366809.21 // v2v receiver initial annual emission
	case "v2v-bx":
		initialTokensPerYear = 17416733.11 // v2v broadcaster initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "V2V client communication mode is not supported!")
	}

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

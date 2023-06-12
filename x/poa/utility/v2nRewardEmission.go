package utility

import (
	"math"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func V2NRewardEmissionPerBlock(ctx sdk.Context, clientCommunicationMode string) (float64, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerMinute := utilConstants.BlocksPerMinute
	currentBlockNumber := int(ctx.BlockHeight())

	var initialTokensPerYear float64
	switch clientCommunicationMode {
	case constants.V2NBX:
		initialTokensPerYear = utilConstants.V2NBXInitialTokenPerYear // v2n broadcaster initial annual emission
	case constants.Runner:
		initialTokensPerYear = utilConstants.RunnerInitialTokenPerYear // v2n receiver (runner) initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2NRewardEmissionPerBlock] failed. V2N communication mode is not supported.")
	}

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / (blocksPerMinute * minutesPerYear)

	// Calculate the number of years that have passed since the start of the token issuance
	yearsSinceStart := int(math.Floor(float64(currentBlockNumber) / (blocksPerMinute * minutesPerYear)))

	// Calculate the number of times the token issuance rate has been halved
	halvings := (yearsSinceStart) / 3

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < halvings; i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / (blocksPerMinute * minutesPerYear)
	}

	return tokensPerBlock, nil
}

func V2NRewardEmissionPerEpoch(ctx sdk.Context, clientCommunicationMode string) (float64, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerMinute := utilConstants.BlocksPerMinute
	blocksPerEpoch := utilConstants.BlocksPerEpoch
	currentBlockNumber := int(ctx.BlockHeight())

	var initialTokensPerYear int
	switch clientCommunicationMode {
	case constants.V2NBX:
		initialTokensPerYear = utilConstants.V2NBXInitialTokenPerYear // v2n broadcaster initial annual emission
	case constants.Runner:
		initialTokensPerYear = utilConstants.RunnerInitialTokenPerYear // v2n receiver (runner) initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2NRewardEmissionPerEpoch] failed. V2N communication mode is not supported.")
	}

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / (blocksPerMinute * minutesPerYear)

	// Calculate the number of years that have passed since the start of the token issuance
	yearsSinceStart := (currentBlockNumber) / (blocksPerMinute * minutesPerYear)

	// Calculate the number of times the token issuance rate has been halved
	halvings := (yearsSinceStart) / 3

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < halvings; i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / (blocksPerMinute * minutesPerYear)
	}

	return float64(tokensPerBlock) * blocksPerEpoch, nil
}

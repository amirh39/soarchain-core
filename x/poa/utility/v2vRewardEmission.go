package utility

import (
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const minutesPerYear = utilConstants.MinutesPerYear

func V2VRewardEmissionPerBlock(ctx sdk.Context, clientCommunicationMode string) (int, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerMinute := utilConstants.V2VRewardEmissionBlockPerYear
	currentBlockNumber := int(ctx.BlockHeight())

	var initialTokensPerYear int
	switch clientCommunicationMode {
	case constants.V2VRX:
		initialTokensPerYear = utilConstants.V2VReceiverInitialAnnual // v2v receiver initial annual emission
	case constants.V2VBX:
		initialTokensPerYear = utilConstants.V2VBroadcasterInitialAnnual // v2v broadcaster initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2VRewardEmissionPerBlock] failed. V2V client communication mode is not supported.")
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

	return tokensPerBlock, nil
}

func V2VRewardEmissionPerEpoch(ctx sdk.Context, clientCommunicationMode string) (float64, error) {

	// Calculates reward coin emissions for each reward type
	blocksPerMinute := utilConstants.V2VRewardEmissionBlockPerYear
	blocksPerEpoch := utilConstants.BlocksPerEpoch
	currentBlockNumber := int(ctx.BlockHeight())

	var initialTokensPerYear int
	switch clientCommunicationMode {
	case constants.V2VRX:
		initialTokensPerYear = utilConstants.V2VReceiverInitialAnnual // v2v receiver initial annual emission
	case constants.V2VBX:
		initialTokensPerYear = utilConstants.V2VBroadcasterInitialAnnual // v2v broadcaster initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2VRewardEmissionPerEpoch] failed. V2V client communication mode is not supported.")
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

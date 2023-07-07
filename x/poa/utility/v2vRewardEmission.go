package utility

import (
	//"math/big"
	"math/big"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility/utilConstants"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const minutesPerYear = utilConstants.MinutesPerYear
const v2vBlockNumbersPerYear = utilConstants.V2VBlocksPerMinute * minutesPerYear

func initialV2VTokensPerYearByClientType(clientCommunicationMode string) (int, error) {
	var initialTokensPerYear int
	switch clientCommunicationMode {
	case constants.V2VRX:
		initialTokensPerYear = utilConstants.V2VReceiverInitialAnnual // v2v receiver initial annual emission
	case constants.V2VBX:
		initialTokensPerYear = utilConstants.V2VBroadcasterInitialAnnual // v2v broadcaster initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2VRewardEmissionPerBlock] failed. V2V client communication mode is not supported.")
	}
	return initialTokensPerYear, nil
}

func v2vHavingsYear(ctx sdk.Context) int {
	currentBlockNumber := big.NewInt(ctx.BlockHeight())

	// Calculate the number of years that have passed since the start of the token issuance
	yearsSinceStart := new(big.Int).Quo(currentBlockNumber, big.NewInt(int64(v2vBlockNumbersPerYear)))

	// Calculate the number of times the token issuance rate has been halved
	halvings := new(big.Int).Quo(yearsSinceStart, big.NewInt(3))
	yearsNumber, _ := strconv.Atoi(halvings.Text(16))
	return yearsNumber
}

func V2VRewardEmissionPerBlock(ctx sdk.Context, clientCommunicationMode string) (int, error) {

	// Calculates reward coin emissions for each reward type

	initialTokensPerYear, err := initialV2VTokensPerYearByClientType(clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2VRewardEmissionPerBlock] failed. V2V client communication mode is not supported.")
	}

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / v2vBlockNumbersPerYear

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < v2vHavingsYear(ctx); i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / v2vBlockNumbersPerYear
	}

	return tokensPerBlock, nil
}

func V2VRewardEmissionPerEpoch(ctx sdk.Context, clientCommunicationMode string) (float64, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerEpoch := utilConstants.BlocksPerEpoch

	initialTokensPerYear, err := initialV2VTokensPerYearByClientType(clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2VRewardEmissionPerBlock] failed. V2V client communication mode is not supported.")
	}

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / v2vBlockNumbersPerYear

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < v2vHavingsYear(ctx); i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / v2vBlockNumbersPerYear
	}

	return float64(tokensPerBlock) * blocksPerEpoch, nil
}

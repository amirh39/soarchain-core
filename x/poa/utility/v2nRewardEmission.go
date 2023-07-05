package utility

import (
	"math/big"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/utility/utilConstants"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const v2nBlockNumbersPerYear = utilConstants.V2NBlocksPerMinute * minutesPerYear

func initialTokensPerYearByClientType(clientCommunicationMode string) (int, error) {
	var initialTokensPerYear int
	switch clientCommunicationMode {
	case constants.V2NBX:
		initialTokensPerYear = utilConstants.V2NBXInitialTokenPerYear // v2n broadcaster initial annual emission
	case constants.Runner:
		initialTokensPerYear = utilConstants.RunnerInitialTokenPerYear // v2n receiver (runner) initial annual emission
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2NRewardEmissionPerBlock] failed. V2N communication mode is not supported.")
	}
	return initialTokensPerYear, nil
}

func havingsYear(ctx sdk.Context) int {
	currentBlockNumber := big.NewInt(ctx.BlockHeight())
	yearsSinceStart := new(big.Int).Quo(currentBlockNumber, big.NewInt(int64(v2nBlockNumbersPerYear)))

	// Calculate the number of times the token issuance rate has been halved
	halvings := new(big.Int).Quo(yearsSinceStart, big.NewInt(3))
	yearsNumber, _ := strconv.Atoi(halvings.Text(16))
	return yearsNumber
}

func V2NRewardEmissionPerBlock(ctx sdk.Context, clientCommunicationMode string) (int, error) {

	// Calculates reward coin emissions for each reward type

	initialTokensPerYear, err := initialTokensPerYearByClientType(clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2NRewardEmissionPerBlock] failed. V2N communication mode is not supported.")
	}

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / v2nBlockNumbersPerYear

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < havingsYear(ctx); i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / v2nBlockNumbersPerYear
	}

	return tokensPerBlock, nil
}

func V2NRewardEmissionPerEpoch(ctx sdk.Context, clientCommunicationMode string) (int, error) {

	// Calculates reward coin emissions for each reward type

	blocksPerEpoch := utilConstants.BlocksPerEpoch

	initialTokensPerYear, err := initialTokensPerYearByClientType(clientCommunicationMode)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[V2NRewardEmissionPerBlock] failed. V2N communication mode is not supported.")
	}

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / v2nBlockNumbersPerYear

	// Update the token issuance rate for each halving that has occurred
	for i := 0; i < havingsYear(ctx); i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / v2nBlockNumbersPerYear
	}

	return tokensPerBlock * blocksPerEpoch, nil
}

package keeper

import (
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {
	epochData, _ := k.GetEpochData(ctx)
	epochCnt := epochData.TotalEpochs
	newEpochCnt := epochCnt + 1

	handleParsingError := func(err error) {
		if err != nil {
			panic(sdkerrors.Wrap(sdkerrors.ErrPanic, "Parsing error"))
		}
	}

	parseUintAndCreateCoin := func(value uint64, multiplier float64) (sdk.Coin, error) {
		amount := float64(value) * multiplier
		return sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(amount))), nil
	}

	var ChallengeCountV2VRx sdk.Coin
	var ChallengeCountV2VBx sdk.Coin
	var ChallengeCountV2NBx sdk.Coin
	var ChallengeCountRunner sdk.Coin
	var ChallengeCountChallenger sdk.Coin

	challengeTypes := []struct {
		field      uint64
		multiplier float64
		target     *sdk.Coin
	}{
		{epochData.V2VRXtotalChallenges, 10.0, &ChallengeCountV2VRx},
		{epochData.V2VBXtotalChallenges, 10.0, &ChallengeCountV2VBx},
		{epochData.V2NBXtotalChallenges, 10.0, &ChallengeCountV2NBx},
		{epochData.RunnerTotalChallenges, 5.0, &ChallengeCountRunner},
		{epochData.ChallengerTotalChallenges, 5.0, &ChallengeCountChallenger},
	}

	for _, c := range challengeTypes {
		challengeCount, err := parseUintAndCreateCoin(c.field, c.multiplier)
		handleParsingError(err)
		*c.target = challengeCount
		mintAndParseCoins(ctx, challengeCount, k)
	}

	newEpochData := types.EpochData{
		TotalEpochs:               newEpochCnt,
		EpochV2VRX:                epochData.EpochV2VRX,
		EpochV2VBX:                epochData.EpochV2VBX,
		EpochV2NBX:                epochData.EpochV2NBX,
		EpochRunner:               epochData.EpochRunner,
		EpochChallenger:           epochData.EpochChallenger,
		V2VRXtotalChallenges:      0,
		V2VBXtotalChallenges:      0,
		V2NBXtotalChallenges:      0,
		RunnerTotalChallenges:     0,
		ChallengerTotalChallenges: 0,
	}

	k.SetEpochData(ctx, newEpochData)
}

func mintAndParseCoins(ctx sdk.Context, coin sdk.Coin, k Keeper) {
	parsedCoin, _ := sdk.ParseCoinNormalized(coin.String())
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{parsedCoin})
}

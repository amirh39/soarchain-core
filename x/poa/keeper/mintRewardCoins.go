package keeper

import (
	"fmt"
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	constant "soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	epochtypes "soarchain/x/epoch/types"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {
	epochData, _ := k.epochKeeper.GetEpochData(ctx)

	if epochData.V2VRXtotalChallenges != 0 || epochData.V2VBXtotalChallenges != 0 || epochData.V2NBXtotalChallenges != 0 || epochData.RunnerTotalChallenges != 0 || epochData.ChallengerTotalChallenges != 0 {
		handleParsingError := func(err error) {
			if err != nil {
				sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Parsing error")
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
			{epochData.V2NBXtotalChallenges, constant.V2NBX, &ChallengeCountV2NBx},
			{epochData.RunnerTotalChallenges, constant.Runner, &ChallengeCountRunner},
			{epochData.ChallengerTotalChallenges, constant.Challenger, &ChallengeCountChallenger},
		}

		for _, c := range challengeTypes {
			challengeCount, err := parseUintAndCreateCoin(c.field, c.multiplier)
			handleParsingError(err)
			*c.target = challengeCount
			mintAndParseCoins(ctx, challengeCount, k)
			// Log the minted coin
			k.Logger(ctx).Info(fmt.Sprintf("Minted coin: %s", challengeCount.String()))
		}

		newEpochData := epochtypes.EpochData{
			TotalEpochs:               epochData.TotalEpochs,
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

		k.epochKeeper.SetEpochData(ctx, newEpochData)
	}
}

func mintAndParseCoins(ctx sdk.Context, coin sdk.Coin, k Keeper) {
	parsedCoin, _ := sdk.ParseCoinNormalized(coin.String())
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{parsedCoin})
}

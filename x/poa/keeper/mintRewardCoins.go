package keeper

import (
	"fmt"
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	constant "soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {
	epochData, _ := k.GetEpochData(ctx)

	if epochData.V2VRXLastBlockChallenges != 0 || epochData.V2VBXLastBlockChallenges != 0 || epochData.V2NBXLastBlockChallenges != 0 || epochData.RunnerLastBlockChallenges != 0 || epochData.ChallengerLastBlockChallenges != 0 {
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
			{epochData.V2VRXLastBlockChallenges, 10.0, &ChallengeCountV2VRx},
			{epochData.V2VBXLastBlockChallenges, 10.0, &ChallengeCountV2VBx},
			{epochData.V2NBXLastBlockChallenges, constant.V2NBX, &ChallengeCountV2NBx},
			{epochData.RunnerLastBlockChallenges, constant.Runner, &ChallengeCountRunner},
			{epochData.ChallengerLastBlockChallenges, constant.Challenger, &ChallengeCountChallenger},
		}

		for _, c := range challengeTypes {
			challengeCount, err := parseUintAndCreateCoin(c.field, c.multiplier)
			handleParsingError(err)
			*c.target = challengeCount
			mintAndParseCoins(ctx, challengeCount, k)
			// Log the minted coin
			k.Logger(ctx).Info(fmt.Sprintf("Minted coin: %s", challengeCount.String()))
		}

		newEpochData := types.EpochData{
			TotalEpochs:                   epochData.TotalEpochs,
			EpochV2VRX:                    epochData.EpochV2VRX,
			EpochV2VBX:                    epochData.EpochV2VBX,
			EpochV2NBX:                    epochData.EpochV2NBX,
			EpochRunner:                   epochData.EpochRunner,
			EpochChallenger:               epochData.EpochChallenger,
			V2VRXtotalChallenges:          epochData.V2VRXtotalChallenges,
			V2VBXtotalChallenges:          epochData.V2VBXtotalChallenges,
			V2NBXtotalChallenges:          epochData.V2NBXtotalChallenges,
			RunnerTotalChallenges:         epochData.RunnerTotalChallenges,
			ChallengerTotalChallenges:     epochData.ChallengerTotalChallenges,
			V2VRXLastBlockChallenges:      0,
			V2VBXLastBlockChallenges:      0,
			V2NBXLastBlockChallenges:      0,
			RunnerLastBlockChallenges:     0,
			ChallengerLastBlockChallenges: 0,
		}

		k.SetEpochData(ctx, newEpochData)
	}
}

func mintAndParseCoins(ctx sdk.Context, coin sdk.Coin, k Keeper) {
	parsedCoin, _ := sdk.ParseCoinNormalized(coin.String())
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{parsedCoin})
}

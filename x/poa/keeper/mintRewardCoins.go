package keeper

import (
	params "soarchain/app/params"
	"soarchain/x/poa/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) MintRewardCoins(ctx sdk.Context) {
	epochData, _ := k.GetEpochData(ctx)
	epochCnt := epochData.TotalEpochs
	newEpochCnt := epochCnt + 1

	parseUintAndCreateCoin := func(value string) (sdk.Coin, error) {
		amount, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return sdk.Coin{}, err
		}
		return sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(amount)), nil
	}

	ChallengeCountV2VRx, err := parseUintAndCreateCoin(epochData.V2VRXtotalChallenges)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrPanic, "Parsing error")
	}

	ChallengeCountV2VBx, err := parseUintAndCreateCoin(epochData.V2VBXtotalChallenges)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrPanic, "Parsing error")
	}

	ChallengeCountV2NBx, err := parseUintAndCreateCoin(epochData.V2NBXtotalChallenges)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrPanic, "Parsing error")
	}

	ChallengeCountRunner, err := parseUintAndCreateCoin(epochData.RunnerTotalChallenges)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrPanic, "Parsing error")
	}

	ChallengeCountChallenger, err := parseUintAndCreateCoin(epochData.ChallengerTotalChallenges)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrPanic, "Parsing error")
	}

	mintAndParseCoins := func(ctx sdk.Context, coin sdk.Coin) {
		parsedCoin, _ := sdk.ParseCoinNormalized(coin.String())
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{parsedCoin})
	}

	mintAndParseCoins(ctx, ChallengeCountV2VRx)
	mintAndParseCoins(ctx, ChallengeCountV2VBx)
	mintAndParseCoins(ctx, ChallengeCountV2NBx)
	mintAndParseCoins(ctx, ChallengeCountRunner)
	mintAndParseCoins(ctx, ChallengeCountChallenger)

	newEpochData := types.EpochData{
		TotalEpochs: newEpochCnt,
		EpochV2VRX:  ChallengeCountV2VRx.String(),
		EpochV2VBX:  ChallengeCountV2VBx.String(),
		EpochV2NBX:  ChallengeCountV2NBx.String(),
		EpochRunner: ChallengeCountRunner.String(),
	}

	k.SetEpochData(ctx, newEpochData)
}

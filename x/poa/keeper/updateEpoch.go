package keeper

import (
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateEpoch(ctx sdk.Context) {
	epochData, _ := k.GetEpochData(ctx)
	epochCnt := epochData.TotalEpochs
	newEpochCnt := epochCnt + 1

	newEpochData := types.EpochData{
		TotalEpochs:               newEpochCnt,
		EpochV2VRX:                sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochV2VBX:                sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochV2NBX:                sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochRunner:               sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		EpochChallenger:           sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		V2VRXtotalChallenges:      0,
		V2VBXtotalChallenges:      0,
		V2NBXtotalChallenges:      0,
		RunnerTotalChallenges:     0,
		ChallengerTotalChallenges: 0,
	}

	k.SetEpochData(ctx, newEpochData)

}
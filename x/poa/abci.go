package poa

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	// EpochDuration defines how long one epoch lasts in minutes
	EpochDuration = 30
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// BeginBlocker for the PoA module. It checks if a new epoch has started and if so,
	// it mints coins from the banking module to the PoA module according to total rewards earned during the epoch.

	// check if a new epoch has started
	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {
		k.MintRewardCoins(ctx)

		// Update epochs passed
		epoch, _ := k.GetEpochData(ctx)
		epochCnt := epoch.TotalEpochs
		newEpochCnt := epochCnt + 1

		newEpoch := types.EpochData{
			TotalEpochs: newEpochCnt,
			EpochV2VRX:  epoch.EpochV2VRX,
			EpochV2VBX:  epoch.EpochV2VBX,
			EpochV2NBX:  epoch.EpochV2NBX,
			EpochRunner: epoch.EpochRunner,
		}

		k.SetEpochData(ctx, newEpoch)
	}
}

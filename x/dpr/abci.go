package dpr

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k AppModule) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	logger := k.keeper.Logger(ctx)

	epochData, found := k.epochKeeper.GetEpochData(ctx)
	if !found {
		logger.Error("Fetching epoch data failed.", "path", "BeginBlocker")
	}

	k.keeper.DeactivateDpr(ctx, epochData.TotalEpochs)

}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

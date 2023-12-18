package dpr

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/amirh39/soarchain-core/x/poa/keeper"
	"github.com/amirh39/soarchain-core/x/poa/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k AppModule) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	logger := k.keeper.Logger(ctx)

	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {

		epochData, found := k.epochKeeper.GetEpochData(ctx)
		if !found {
			logger.Error("[Dpr Module][BeginBlocker] Fetching epoch data failed.", "path", "BeginBlocker")
		}
		k.keeper.DeactivateDpr(ctx, epochData.TotalEpochs)

		//k.keeper.DistributeRewards(ctx)
	}
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

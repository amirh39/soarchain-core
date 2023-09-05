package did

import (
	"log"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"soarchain/x/did/keeper"
	"soarchain/x/did/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k AppModule) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	logger := k.keeper.Logger(ctx)

	log.Println("############## Begin Blocker Started ##############")

	// BeginBlocker for the did module. It checks if a new epoch has started and if so,
	// it mints coins from the banking module to the did module according to total rewards earned during the epoch.

	// check if a new epoch has started
	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {
		if logger != nil {
			logger.Info("Update epoch started.", "path", "BeginBlocker")
		}

		if logger != nil {
			logger.Info("Fetching epoch data successfully done.")
		}

		if logger != nil {
			logger.Info("Update epoch  successfully done.")
		}

	}

	log.Println("############## End of Begin Blocker ##############")
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}

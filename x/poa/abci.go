package poa

import (
	"log"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k AppModule) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	epochData, found := k.epochKeeper.GetEpochData(ctx)

	logger := k.keeper.Logger(ctx)

	log.Println("############## Begin Blocker Started ##############")

	// BeginBlocker for the PoA module. It checks if a new epoch has started and if so,
	// it mints coins from the banking module to the PoA module according to total rewards earned during the epoch.

	err := k.keeper.MintRewardCoins(ctx, epochData)
	if err != nil {
		logger.Info("Mint Reward Coins failed.", "path", "BeginBlocker")
	}

	// check if a new epoch has started
	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {

		k.epochKeeper.UpdateEpoch(ctx)

		if logger != nil {
			logger.Info("Update epoch  successfully done.", "path", "BeginBlocker", "epoch data", epochData, "found", found)
		}
	}

	if (epochData.TotalEpochs%192 == 0) && (epochData.TotalEpochs != 0) {

		epochData, err = k.keeper.ComputeAdaptiveHalving(ctx, epochData)
		if logger != nil && err == nil {
			logger.Info("Compute Adaptive Halving successfully done.", "path", "BeginBlocker")
		}

		k.epochKeeper.SetEpochData(ctx, epochData)
	}

	log.Println("############## End of Begin Blocker ##############")
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

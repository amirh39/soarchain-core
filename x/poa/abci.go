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

	logger := k.keeper.Logger(ctx)

	log.Println("############## Begin Blocker Started ##############")

	// BeginBlocker for the PoA module. It checks if a new epoch has started and if so,
	// it mints coins from the banking module to the PoA module according to total rewards earned during the epoch.

	if logger != nil {
		logger.Info("Mint Reward Coins started.", "path", "BeginBlocker")
	}

	k.keeper.MintRewardCoins(ctx)

	if logger != nil {
		logger.Info("Mint Reward Coins successfully done.", "path", "BeginBlocker")
	}

	// check if a new epoch has started
	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {
		if logger != nil {
			logger.Info("Update epoch started.", "path", "BeginBlocker")
		}

		item, found := k.epochKeeper.GetEpochData(ctx)

		if logger != nil {
			logger.Info("Update epoch  successfully done.", "path", "BeginBlocker", "epoch data", item, "found", found)
		}

		k.epochKeeper.UpdateEpoch(ctx)

	}

	epochData, _ := k.epochKeeper.GetEpochData(ctx)
	if (epochData.TotalEpochs%192 == 0) && (epochData.TotalEpochs != 0) {
		log.Println("ComputeAdaptiveHalving")
		//k.ComputeAdaptiveHalving(ctx)

	}
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}

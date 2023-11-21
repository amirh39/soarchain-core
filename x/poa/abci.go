package poa

import (
	"log"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soar-robotics/soarchain-core/x/poa/keeper"
	"github.com/soar-robotics/soarchain-core/x/poa/types"

	epochtypes "github.com/soar-robotics/soarchain-core/x/epoch/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k AppModule) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	epochData, found := k.epochKeeper.GetEpochData(ctx)

	logger := k.keeper.Logger(ctx)

	log.Println("############## Poa Module Begin Blocker is Started ##############")

	// BeginBlocker for the PoA module. It checks if a new epoch has started and if so,
	// it mints coins from the banking module to the PoA module according to total rewards earned during the epoch.

	err := k.keeper.MintRewardCoins(ctx, epochData)
	if err != nil {
		logger.Info("[Poa Module][BeginBlocker] Mint Reward Coins failed.", "path", "BeginBlocker")
	}

	// check if a new epoch has started
	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {

		// apply halving
		if epochData.TotalEpochs%192 == 0 {
			epochData, err = k.keeper.ComputeAdaptiveHalving(ctx, epochData)
			if err != nil {
				// Log the error
				if logger != nil {
					logger.Error("[Poa Module][BeginBlocker] Error in Compute Adaptive Halving.", "path", "BeginBlocker", "error", err)
				}

			} else if logger != nil {
				// If there's no error, log the success message
				logger.Info("[Poa Module][BeginBlocker] Compute Adaptive Halving successfully done.", "path", "BeginBlocker")
			}

			k.epochKeeper.SetEpochData(ctx, epochData)
		}
		k.epochKeeper.UpdateEpoch(ctx)
		if logger != nil {
			logger.Info("[Poa Module][BeginBlocker] Update epoch successfully done.", "path", "BeginBlocker", "epoch data", epochData, "found", found)
		}

		randomNumber, found := k.epochKeeper.RandomNumber(ctx, epochData.TotalEpochs)
		if !found {
			logger.Error("[Poa Module][BeginBlocker] Fetching random number failed.", "path", "BeginBlocker")
		}

		randomData := epochtypes.RandomData{
			Id:           strconv.FormatInt(ctx.BlockHeight(), 10),
			EpochNumber:  strconv.FormatUint(epochData.TotalEpochs, 10),
			RandomNumber: randomNumber,
		}

		k.epochKeeper.SetRandomData(ctx, randomData)
		if logger != nil {
			logger.Info("[Poa Module][BeginBlocker] Update random number successfully done.", "path", "BeginBlocker")
		}
	}
	log.Println("############## End of Poa Module Begin Blocker ##############")
}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}

package poa

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// BeginBlocker for the PoA module. It checks if a new epoch has started and if so,
	// it mints coins from the banking module to the PoA module according to total rewards earned during the epoch.

	k.MintRewardCoins(ctx)
	// check if a new epoch has started
	if (ctx.BlockHeight()%30 == 0) && (ctx.BlockHeight() != 0) {

	}

}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}

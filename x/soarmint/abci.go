package soarmint

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/amirh39/soarchain-core/x/soarmint/keeper"
	"github.com/amirh39/soarchain-core/x/soarmint/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch minter from KV store
	minter, _ := k.GetMinter(ctx)

	// staking rewards ended (minter.Inflation = 0)
	if minter.Inflation.Equal(sdk.ZeroDec()) {
		return
	}

	// fetch stored params
	params := k.GetParams(ctx)

	currentBlock := uint64(ctx.BlockHeight())

	// fetch current token supply
	totalSupply := k.TokenSupply(ctx, params.MintDenom)
	if totalSupply == sdk.ZeroInt() { // debug
		return
	}

	nextInf := minter.PhaseInflationRate(minter.Phase)
	nextPhase := minter.NextPhase(ctx, params)

	mintedCoins := minter.StakingRewardsPerBlock(ctx, params)

	newMinter := types.Minter{
		Inflation:        nextInf,
		Phase:            nextPhase,
		StartPhaseBlock:  currentBlock,
		AnnualProvisions: minter.NextAnnualProvisions(params, totalSupply),
		TargetSupply:     mintedCoins.AmountOf(params.MintDenom),
	}
	k.SetMinter(ctx, newMinter)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// if mintedCoin.Amount.IsInt64() {
	// 	defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	// }

	// ctx.EventManager().EmitEvent(
	// 	sdk.NewEvent(
	// 		types.EventTypeMint,
	// 		sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
	// 		sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
	// 		sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
	// 	),
	// )

}

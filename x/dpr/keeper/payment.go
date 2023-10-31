package keeper

import (
	"fmt"
	"math/big"
	"soarchain/app/params"
	"soarchain/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DistributeRewards(ctx sdk.Context) {
	logger := k.Logger(ctx)
	activeDprs := k.GetAllActiveDpr(ctx)

	if activeDprs == nil {
		logger.Info("There is no active DPR", "path", "DistributeRewards")
		return
	}

	for _, dpr := range activeDprs {
		budget, _ := sdk.ParseCoinNormalized(dpr.DprBudget)

		// Convert budget.Amount to big.Int for precise calculations
		budgetAmountBigInt := big.NewInt(budget.Amount.Int64())

		// Calculate reward per epoch using big.Int
		rewardPerEpochBigInt := new(big.Int).Div(budgetAmountBigInt, big.NewInt(int64(dpr.Duration)))

		// Calculate the reward per client using big.Int
		rewardPerClientBigInt := new(big.Int).Div(rewardPerEpochBigInt, big.NewInt(int64(dpr.MaxClientCount)))

		for _, clientPubKey := range dpr.ClientPubkeys {
			// Convert the rewardPerClientBigInt back to sdk.Int for creating sdk.Coin
			rewardPerClientAmount := sdk.NewIntFromBigInt(rewardPerClientBigInt)

			// Distribute 'rewardPerClient' to the client identified by 'clientPubKey'
			rewardPerClient := sdk.NewCoin(budget.Denom, rewardPerClientAmount)
			client, _ := k.poaKeeper.GetReputation(ctx, clientPubKey)
			netEarnings, _ := sdk.ParseCoinNormalized(client.DprEarnings)

			totalEarnings := netEarnings.Add(rewardPerClient)

			client.DprEarnings = totalEarnings.String()

			k.poaKeeper.SetReputation(ctx, client)

			// Example print statement (replace with actual distribution logic)
			fmt.Printf("Distributing %s to client with pubkey: %s\n", rewardPerClient, clientPubKey)
		}
	}

	if logger != nil {
		logger.Info("Payments of the DPR successfully Done.", "path", "DistributeRewards")
	}

}

func (k Keeper) calculateTotalEarnings(ctx sdk.Context, currentEarnings string, earnedRewardsBigInt *big.Int, entityType string) (sdk.Coin, error) {
	earnedAmount := sdk.NewIntFromBigInt(earnedRewardsBigInt)
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedAmount)

	netEarnings, err := sdk.ParseCoinNormalized(currentEarnings)
	if err != nil {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errors.NetEarnings)
	}

	totalEarnings := netEarnings.Add(earnedCoin)

	return totalEarnings, nil
}

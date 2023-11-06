package keeper

import (
	"log"
	"soarchain/app/params"
	"soarchain/x/did/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DistributeRewards(ctx sdk.Context, clientDid types.ClientDid) (sdk.DecCoins, error) {
	// Check if there are any DprInfos in the client DID
	if len(clientDid.GetDprInfos()) == 0 {
		// If there are no DprInfos, exit the function.
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[DistributeRewards] No DPRs found in Client DID")
	}

	var totalRewards sdk.DecCoins

	// Get current epoch from the context or state
	currentEpoch, _ := k.epochKeeper.GetEpochData(ctx)

	// Iterate over the DPRs the client is in
	for i, dprInfo := range clientDid.DprInfos {
		// Retrieve DPR using dprInfo.Id
		dpr, _ := k.GetDpr(ctx, dprInfo.Id)

		// Calculate total reward possible until the current epoch
		elapsedEpochs := currentEpoch.TotalEpochs - dpr.DprStartEpoch
		if elapsedEpochs > dpr.Duration {
			elapsedEpochs = dpr.Duration // Cap it at the duration
		}
		log.Println(dpr.DprBudget)

		dprBudgetCoin, err := sdk.ParseCoinNormalized(dpr.DprBudget)
		if err != nil {
			log.Printf("Error parsing DPR budget: %v\n", err)
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[DistributeRewards] Error parsing DPR budget.")
		}

		dprBudget, err := strconv.ParseInt(dprBudgetCoin.Amount.String(), 10, 64)
		if err != nil {
			log.Printf("Error converting DPR budget amount to int: %v\n", err)
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[DistributeRewards] Error converting DPR budget amount.")
		}

		dprBudgetDec := sdk.NewDec(dprBudget)
		log.Println(dprBudgetDec)
		rewardPerEpoch := sdk.NewDecCoinFromDec(params.BondDenom, dprBudgetDec.Quo(sdk.NewDec(int64(dpr.MaxClientCount))).Quo(sdk.NewDec(int64(dpr.Duration))))
		log.Println(rewardPerEpoch)
		if rewardPerEpoch.IsNegative() || rewardPerEpoch.IsZero() {
			log.Printf("Reward per epoch is not positive: %s\n", rewardPerEpoch)
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[DistributeRewards] Reward per epoch is not positive.")
		}

		totalPossibleRewardDec := rewardPerEpoch.Amount.MulInt64(int64(elapsedEpochs))
		log.Println(totalPossibleRewardDec)
		if totalPossibleRewardDec.IsNegative() {
			log.Printf("Total possible reward is negative: %s\n", totalPossibleRewardDec)
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[DistributeRewards] Total possible reward is negative.")
		}
		totalPossibleReward := sdk.NewDecCoinFromDec(params.BondDenom, totalPossibleRewardDec)
		log.Println(totalPossibleReward)
		if !totalPossibleReward.IsValid() {
			log.Printf("Total possible reward is invalid: %s\n", totalPossibleReward)
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[DistributeRewards] Total possible reward is invalid.")
		}
		// Subtract the amount already claimed
		claimedAmount, ok := sdk.NewIntFromString(dprInfo.Claimed)
		log.Println(claimedAmount)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[DistributeRewards] Invalid claimed amount")
		}
		rewardToDistribute := totalPossibleReward.Sub((sdk.NewDecCoin(params.BondDenom, claimedAmount)))
		log.Println(rewardToDistribute)
		currentClaimedAmount, ok := sdk.NewIntFromString(dprInfo.Claimed)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[DistributeRewards] Invalid claimed amount")
		}

		// Calculate the new claimed amount by adding the rewardToDistribute to the current claimed amount
		newClaimedAmount := currentClaimedAmount.Add(rewardToDistribute.Amount.RoundInt())
		log.Println(newClaimedAmount)
		// Update the claimed amount in the specific DprInfo object
		clientDid.DprInfos[i].Claimed = newClaimedAmount.String()
		k.didKeeper.SetClientDid(ctx, clientDid)
		// Add the calculated reward to the total rewards
		totalRewards = totalRewards.Add(rewardToDistribute)

	}

	// Now totalRewards contains the amount to be distributed to be used in the transaction flow
	return totalRewards, nil
}

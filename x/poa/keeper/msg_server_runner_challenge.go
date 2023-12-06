package keeper

import (
	"context"
	"log"
	"math/big"
	params "soarchain/app/params"
	epoch "soarchain/x/epoch/types"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/errors"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) updateChallengerReputation(ctx sdk.Context, challengerReputation types.Reputation, epoch epoch.EpochData) error {

	var totalEarnings sdk.Coin
	var rewardMultiplier float64
	newScore := make([]float64, 0)
	rewardMultiplier, score := k.rewardAndScore(challengerReputation.Score)
	newScore = append(newScore, score)

	totalAmount := big.NewInt(int64(epoch.ChallengerPerChallengeValue))
	earnedRewardsBigInt, err := utility.CalculateRewards(totalAmount, newScore, []int{1})
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, errors.EarnedRewardsBigInt)
	}

	if len(earnedRewardsBigInt) > 0 {
		totalEarnings, err = k.calculateTotalEarnings(ctx, challengerReputation.NetEarnings, earnedRewardsBigInt[0], constants.Challenger)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.TotalEarnings)
		}
	}

	updatedChallenger := types.Reputation{
		PubKey:             challengerReputation.PubKey,
		Address:            challengerReputation.Address,
		Score:              strconv.FormatFloat(newScore[0], 'f', -1, 64),
		NetEarnings:        totalEarnings.String(),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, challengerReputation.Address), 10),
		Type:               challengerReputation.Type,
		StakedAmount:       challengerReputation.StakedAmount,
		DprEarnings:        challengerReputation.DprEarnings,
	}

	k.SetReputation(ctx, updatedChallenger)

	return nil
}

func (k Keeper) coolDownMultiplier(ctx sdk.Context, creator string) uint64 {

	vrfData, _ := k.CreateVRF(ctx, creator, constants.Multiplier)

	generatedNumber, _ := strconv.ParseUint(vrfData.FinalVrv, 10, 64)

	var coolDownMultiplier uint64
	if generatedNumber > 0 {
		coolDownMultiplier = generatedNumber
	} else {
		coolDownMultiplier = 1
	}

	return coolDownMultiplier
}

func (k Keeper) rewardAndScore(score string) (float64, float64) {
	scoreFloat64, _ := strconv.ParseFloat(score, 64)
	newScore := utility.CalculateScore(scoreFloat64, true)
	rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

	return rewardMultiplier, newScore
}

func (k Keeper) punish(score string) (float64, float64) {
	scoreFloat64, _ := strconv.ParseFloat(score, 64)
	newScore := utility.CalculateScore(scoreFloat64, false)
	rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

	return rewardMultiplier, newScore
}

func (k Keeper) updateRunnerReputation(ctx sdk.Context, creator string, runnerPubKey string, result string, epoch epoch.EpochData) error {
	runner, found := k.GetReputation(ctx, runnerPubKey)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.NotFoundAValidRunner)
	}

	var totalEarnings sdk.Coin
	var score float64
	newScore := make([]float64, 0)
	var rewardMultiplier float64

	if result == constants.Reward {
		rewardMultiplier, score = k.rewardAndScore(runner.Score)
	} else if result == constants.Punish {
		rewardMultiplier, score = k.punish(runner.Score)
	} else {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.InvaldChallengeResult)
	}
	newScore = append(newScore, score)

	earnedRewardsBigInt, err := utility.CalculateRewards(big.NewInt(int64(epoch.RunnerPerChallengeValue)), newScore, []int{1})
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, errors.EarnedRewardsBigInt)
	}
	if len(earnedRewardsBigInt) > 0 {
		totalEarnings, err = k.calculateTotalEarnings(ctx, runner.NetEarnings, earnedRewardsBigInt[0], constants.Runner)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.TotalEarnings)
		}
	}

	updatedRunner := types.Reputation{
		PubKey:             runner.PubKey,
		Address:            runner.Address,
		Score:              strconv.FormatFloat(newScore[0], 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        totalEarnings.String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, creator), 10),
		Type:               runner.Type,
		StakedAmount:       runner.StakedAmount,
		DprEarnings:        runner.DprEarnings,
	}
	k.SetReputation(ctx, updatedRunner)

	return nil
}

func (k Keeper) updateReputation(ctx sdk.Context, msg *types.MsgRunnerChallenge, epoch epoch.EpochData) error {
	clientPubkeysCount := len(msg.Clients)
	if clientPubkeysCount < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.NoV2nBxAddressPubKeys)
	}

	// Create arrays to store scores and message counts
	scores := make([]float64, clientPubkeysCount)
	messageCounts := make([]int, clientPubkeysCount)

	for i := 0; i < clientPubkeysCount; i++ {
		reputation, isFound := k.GetReputation(ctx, msg.Clients[i].P)
		if !isFound {
			return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, errors.NotFoundAClient)
		}

		score, err := strconv.ParseFloat(reputation.Score, 64)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid score")
		}

		// Set score and message count in corresponding arrays
		scores[i] = score
		messageCounts[i] = int(msg.Clients[i].N)
	}

	// Now, you have arrays of scores and message counts that correspond to each other

	// Calculate rewards for all scores
	earnedRewardsBigInt, err := utility.CalculateRewards(big.NewInt(int64(epoch.V2NBXPerChallengeValue)), scores, messageCounts)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, errors.EarnedRewardsBigInt)
	}
	var totalEarnings sdk.Coin
	for i := 0; i < clientPubkeysCount; i++ {
		reputation, isFound := k.GetReputation(ctx, msg.Clients[i].P)
		if !isFound {
			return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, errors.NotFoundAClient)
		}
		// RewardAndScore functionality
		rewardMultiplier, score := k.rewardAndScore(reputation.Score)

		if len(earnedRewardsBigInt) > 0 {
			totalEarnings, err = k.calculateTotalEarnings(ctx, reputation.NetEarnings, earnedRewardsBigInt[i], constants.V2NBX)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.TotalEarnings)
			}
		}

		updatedReputation := types.Reputation{
			PubKey:             reputation.PubKey,
			Address:            reputation.Address,
			Score:              strconv.FormatFloat(score, 'f', -1, 64),
			NetEarnings:        totalEarnings.String(),
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, msg.Creator), 10),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			Type:               reputation.Type,
			StakedAmount:       reputation.StakedAmount,
			DprEarnings:        reputation.DprEarnings,
		}

		k.SetReputation(ctx, updatedReputation)
	}
	return nil
}

func (k Keeper) calculateTotalEarnings(ctx sdk.Context, currentEarnings string, earnedRewardsBigInt *big.Int, entityType string) (sdk.Coin, error) {
	earnedAmount := sdk.NewIntFromBigInt(earnedRewardsBigInt)
	earnedCoin := sdk.NewCoin(params.BondDenom, earnedAmount)

	netEarnings, err := sdk.ParseCoinNormalized(currentEarnings)
	if err != nil {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errors.NetEarnings)
	}

	totalEarnings := netEarnings.Add(earnedCoin)

	// Update the epoch rewards
	if epochErr := k.epochKeeper.UpdateEpochRewards(ctx, entityType, earnedCoin); epochErr != nil {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidType, errors.EpochError)
	}

	return totalEarnings, nil
}

func (k msgServer) RunnerChallenge(goCtx context.Context, msg *types.MsgRunnerChallenge) (*types.MsgRunnerChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)
	log.Println("############## Runner challenge Transaction Has Started ##############")

	epochData, isFound := k.epochKeeper.GetEpochData(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RunnerChallenge][GetEpochData] failed. Epoch data is not found!")
	}

	challengerReputation, found := k.GetReputationsByAddressAndType(ctx, msg.Creator, constants.V2NChallenger)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.GetChallengerByType)
	}

	err := k.updateReputation(ctx, msg, epochData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	}

	if logger != nil {
		logger.Info("Updating reputation successfully done.", "transaction", "RunnerChallenge")
	}

	err = k.updateRunnerReputation(ctx, msg.Creator, msg.Runner, msg.Result, epochData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	}

	if logger != nil {
		logger.Info("Updating runner successfully done.", "transaction", "RunnerChallenge")
	}

	/** Update challenger info after the successfull reward session */
	err = k.updateChallengerReputation(ctx, challengerReputation, epochData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	}

	if logger != nil {
		logger.Info("Updating challengerReputation successfully done.", "transaction", "RunnerChallenge")
	}

	//update the challenge counts
	epochErr := k.epochKeeper.UpdateEpochRewards(ctx, "runner_challenge", sdk.NewCoin(params.BondDenom, sdk.ZeroInt()))
	if epochErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EpochError)
	}

	if logger != nil {
		logger.Info("Updating epoch reward successfully done.", "transaction", "RunnerChallenge")
	}

	log.Println("############## End of Runner Challenge Transaction ##############")

	return &types.MsgRunnerChallengeResponse{}, nil
}

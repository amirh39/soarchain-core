package keeper

import (
	"context"
	"log"
	"math/big"
	param "soarchain/app/params"
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

func (k Keeper) updateChallenger(ctx sdk.Context, challenger types.Challenger, epoch epoch.EpochData) error {

	var totalEarnings sdk.Coin
	var rewardMultiplier float64
	newScore := make([]float64, 0)
	rewardMultiplier, score := k.rewardAndScore(challenger.Score)
	newScore = append(newScore, score)

	totalAmount := big.NewInt(int64(epoch.ChallengerPerChallengeValue))
	earnedRewardsBigInt, err := utility.CalculateRewards(totalAmount, newScore)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, errors.EarnedRewardsBigInt)
	}

	if len(earnedRewardsBigInt) > 0 {
		earnedAmount := sdk.NewIntFromBigInt(earnedRewardsBigInt[0])

		earnedCoin := sdk.NewCoin(param.BondDenom, earnedAmount)

		netEarnings, err := sdk.ParseCoinNormalized(challenger.NetEarnings)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errors.NetEarnings)
		}

		totalEarnings = netEarnings.Add(earnedCoin)

		if epochErr := k.epochKeeper.UpdateEpochRewards(ctx, constants.Challenger, earnedCoin); epochErr != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidType, errors.EpochError)
		}
	}

	updatedChallenger := types.Challenger{
		PubKey:           challenger.PubKey,
		Address:          challenger.Address,
		Score:            strconv.FormatFloat(newScore[0], 'f', -1, 64),
		StakedAmount:     challenger.StakedAmount,
		NetEarnings:      totalEarnings.String(),
		Type:             challenger.Type,
		IpAddress:        challenger.IpAddress,
		RewardMultiplier: strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
	}

	k.SetChallenger(ctx, updatedChallenger)

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

func (k Keeper) updateRunner(ctx sdk.Context, creator string, runnerPubKey string, result string, epoch epoch.EpochData) error {
	runner, found := k.GetRunnerUsingPubKey(ctx, runnerPubKey)
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

	earnedRewardsBigInt, err := utility.CalculateRewards(big.NewInt(int64(epoch.RunnerPerChallengeValue)), newScore)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, errors.EarnedRewardsBigInt)
	}

	if len(earnedRewardsBigInt) > 0 {
		earnedAmount := sdk.NewIntFromBigInt(earnedRewardsBigInt[0])

		earnedCoin := sdk.NewCoin(param.BondDenom, earnedAmount)

		netEarnings, err := sdk.ParseCoinNormalized(runner.NetEarnings)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errors.NetEarnings)
		}

		totalEarnings = netEarnings.Add(earnedCoin)

		// Update the epoch rewards
		if epochErr := k.epochKeeper.UpdateEpochRewards(ctx, constants.Runner, earnedCoin); epochErr != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidType, errors.EpochError)
		}
	}

	updatedRunner := types.Runner{
		PubKey:             runner.PubKey,
		Address:            runner.Address,
		Score:              strconv.FormatFloat(newScore[0], 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		StakedAmount:       runner.StakedAmount,
		NetEarnings:        totalEarnings.String(),
		IpAddress:          runner.IpAddress,
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, creator), 10),
	}
	k.SetRunner(ctx, updatedRunner)

	return nil
}

func (k Keeper) updateClient(ctx sdk.Context, msg *types.MsgRunnerChallenge, epoch epoch.EpochData) error {
	v2nBxAddrCount := len(msg.ClientPubkeys)
	if v2nBxAddrCount < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.NoV2nBxAddressPubKeys)
	}

	// Create an array of scores to send to CalculateRewards
	scores := make([]float64, v2nBxAddrCount)
	for i := 0; i < v2nBxAddrCount; i++ {
		v2nBxClient, isFound := k.GetClient(ctx, msg.ClientPubkeys[i])
		if !isFound {
			return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, errors.NotFoundAClient)
		}

		score, err := strconv.ParseFloat(v2nBxClient.Score, 64)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid score")
		}
		scores[i] = score
	}

	// Calculate rewards for all scores
	rewards, err := utility.CalculateRewards(big.NewInt(int64(epoch.V2NBXPerChallengeValue)), scores)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, errors.EarnedRewardsBigInt)
	}

	for i := 0; i < v2nBxAddrCount; i++ {
		v2nBxClient, isFound := k.GetClient(ctx, msg.ClientPubkeys[i])
		if !isFound {
			return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, errors.NotFoundAClient)
		}

		// RewardAndScore functionality
		rewardMultiplier, score := k.rewardAndScore(v2nBxClient.Score)

		earnedAmount := sdk.NewIntFromBigInt(rewards[i])

		earnedCoin := sdk.NewCoin(param.BondDenom, earnedAmount)

		if epochErr := k.epochKeeper.UpdateEpochRewards(ctx, constants.V2NBX, earnedCoin); epochErr != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidType, errors.EpochError)
		}

		netEarnings, err := sdk.ParseCoinNormalized(v2nBxClient.NetEarnings)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errors.NetEarnings)
		}

		totalEarnings := netEarnings.Add(earnedCoin)

		updatedClient := types.Client{
			Index:              v2nBxClient.Index,
			Address:            v2nBxClient.Address,
			Score:              strconv.FormatFloat(score, 'f', -1, 64),
			NetEarnings:        totalEarnings.String(),
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, msg.Creator), 10),
			Type:               v2nBxClient.Type,
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		}

		k.SetClient(ctx, updatedClient)

		k.updateMotusWallet(ctx, v2nBxClient.Address, updatedClient)
	}

	return nil
}

func (k Keeper) updateMotusWallet(ctx sdk.Context, address string, client types.Client) {
	motusWallet, _ := k.GetMotusWallet(ctx, address)

	newMotusWallet := types.MotusWallet{
		Index:  motusWallet.Index,
		Client: &client,
	}

	k.SetMotusWallet(ctx, newMotusWallet)
}

func (k msgServer) RunnerChallenge(goCtx context.Context, msg *types.MsgRunnerChallenge) (*types.MsgRunnerChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)
	log.Println("############## runner_challenge Transaction Has Started ##############")

	epochData, isFound := k.epochKeeper.GetEpochData(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RunnerChallenge][GetEpochData] failed. Epoch data is not found!")
	}

	challenger, found := k.GetChallengerByType(ctx, msg.Creator, constants.V2NChallengerType)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.GetChallengerByType)
	}

	err := k.updateRunner(ctx, msg.Creator, msg.RunnerPubkey, msg.ChallengeResult, epochData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	}

	if logger != nil {
		logger.Info("Updating runner successfully done.", "transaction", "RunnerChallenge")
	}

	err = k.updateClient(ctx, msg, epochData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	}

	if logger != nil {
		logger.Info("Updating client successfully done.", "transaction", "RunnerChallenge")
	}

	/** Update challenger info after the successfull reward session */
	err = k.updateChallenger(ctx, challenger, epochData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	}

	if logger != nil {
		logger.Info("Updating challenger successfully done.", "transaction", "RunnerChallenge")
	}

	//update the challenge counts
	if epochErr := k.epochKeeper.UpdateEpochRewards(ctx, "runner_challenge", sdk.NewCoin(params.BondDenom, sdk.ZeroInt())); epochErr == nil {
		log.Println("UpdateEpochRewards working")
	} else {
		// There was an error in the UpdateEpochRewards function, handle it here
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EpochError)
	}
	log.Println("############## End of Runner Challenge Transaction ##############")

	return &types.MsgRunnerChallengeResponse{}, nil
}

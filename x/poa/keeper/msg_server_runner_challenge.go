package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	param "soarchain/app/params"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/errors"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
)

func (k Keeper) updateChallenger(ctx sdk.Context, challenger types.Challenger) {

	var totalEarnings sdk.Coin
	var rewardMultiplier float64
	newScore := make([]float64, 0)
	rewardMultiplier, score := k.rewardAndScore(challenger.Score)
	newScore = append(newScore, score)

	earnedRewardsFloat := k.CalculateRewards(5, newScore)

	if len(earnedRewardsFloat) > 0 {
		earnedRewardsInt := sdk.NewIntFromUint64(uint64(earnedRewardsFloat[0]))
		earnedCoin := sdk.NewCoin(param.BondDenom, earnedRewardsInt)

		netEarnings, err := sdk.ParseCoinNormalized(challenger.NetEarnings)
		if err != nil {
			// Handle the error appropriately
		}

		totalEarnings = netEarnings.Add(earnedCoin)

		if epochErr := k.UpdateEpochRewards(ctx, "challenger", earnedCoin); epochErr != nil {
			// Handle the error appropriately
		}
	}

	updatedChallenger := types.Challenger{
		PubKey:           challenger.PubKey,
		Address:          challenger.Address,
		Score:            strconv.FormatFloat(newScore[0], 'f', -1, 64),
		StakedAmount:     challenger.StakedAmount,
		NetEarnings:      totalEarnings.String(),
		Type:             challenger.Type,
		IpAddr:           challenger.IpAddr,
		RewardMultiplier: strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
	}

	k.SetChallenger(ctx, updatedChallenger)
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

func (k Keeper) updateRunner(ctx sdk.Context, creator string, runnerPubKey string, result string) error {

	runner, found := k.GetRunnerUsingPubKey(ctx, runnerPubKey)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.NotFoundAValidRunner)
	}

	var totalEarnings sdk.Coin
	var rewardMultiplier float64
	var score float64
	newScore := make([]float64, 0)

	if result == constants.Reward {
		rewardMultiplier, score = k.rewardAndScore(runner.Score)
	} else if result == constants.Punish {
		rewardMultiplier, score = k.punish(runner.Score)
	} else {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.InvaldChallengeResult)
	}
	newScore = append(newScore, score)
	earnedRewardsFloat := k.CalculateRewards(5, newScore)

	if len(earnedRewardsFloat) > 0 {
		earnedRewardsInt := sdk.NewIntFromUint64(uint64(earnedRewardsFloat[0]))
		earnedCoin := sdk.NewCoin(param.BondDenom, earnedRewardsInt)

		netEarnings, err := sdk.ParseCoinNormalized(runner.NetEarnings)
		if err != nil {
			// Handle the error appropriately
		}

		totalEarnings = netEarnings.Add(earnedCoin)

		if epochErr := k.UpdateEpochRewards(ctx, "challenger", earnedCoin); epochErr != nil {
			// Handle the error appropriately
		}
	}

	updatedRunner := types.Runner{
		PubKey:             runner.PubKey,
		Address:            runner.Address,
		Score:              strconv.FormatFloat(newScore[0], 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		StakedAmount:       runner.StakedAmount,
		NetEarnings:        totalEarnings.String(),
		IpAddr:             runner.IpAddr,
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, creator), 10),
	}
	k.SetRunner(ctx, updatedRunner)

	return nil
}

// func (k Keeper) updateClient(ctx sdk.Context, msg *types.MsgRunnerChallenge) error {

// 	// ToDo: Set MOTUS mini rewards
// 	v2nBxAddrCount := len(msg.ClientPubkeys)
// 	if v2nBxAddrCount < 1 {
// 		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.NoV2nBxAddrPubKeys)
// 	}

// 	/** All MOTUS mini devices will be rewarded */
// 	for i := 0; i < v2nBxAddrCount; i++ {

// 		v2nBxClient, isFound := k.GetClient(ctx, msg.ClientPubkeys[i])
// 		if !isFound {
// 			return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, errors.NotFoundAClient)
// 		}

// 		rewardMultiplier, newScore := k.rewardAndScore(v2nBxClient.Score)

// 		totalEarnings, err := k.totalEarnings(ctx, v2nBxClient.NetEarnings, constants.V2NBX)
// 		if err != nil {
// 			return sdkerrors.Wrap(sdkerrors.ErrPanic, errors.TotalEarnings)
// 		}

// 		updatedClient := types.Client{
// 			Index:              v2nBxClient.Index,
// 			Address:            v2nBxClient.Address,
// 			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
// 			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
// 			NetEarnings:        totalEarnings.String(),
// 			LastTimeChallenged: ctx.BlockTime().String(),
// 			CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, msg.Creator), 10),
// 			Type:               v2nBxClient.Type,
// 		}

// 		k.SetClient(ctx, updatedClient)

// 		k.updateMotusWallet(ctx, v2nBxClient.Address, updatedClient)

// 	}

// 	return nil
// }

// func (k Keeper) updateMotusWallet(ctx sdk.Context, address string, client types.Client) {

// 	motusWallet, _ := k.GetMotusWallet(ctx, address)

// 	newMotusWallet := types.MotusWallet{
// 		Index:  motusWallet.Index,
// 		Client: &client,
// 	}

// 	k.SetMotusWallet(ctx, newMotusWallet)
// }

func (k msgServer) RunnerChallenge(goCtx context.Context, msg *types.MsgRunnerChallenge) (*types.MsgRunnerChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, found := k.GetChallengerByType(ctx, msg.Creator, constants.V2NChallengerType)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.GetChallengerByType)
	}

	// err := k.updateRunner(ctx, msg.Creator, msg.RunnerpubKey, msg.ChallengeResult)
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	// }

	// err = k.updateClient(ctx, msg)
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errors.EarnedTokenRewardsFloat)
	// }

	/** Update challenger info after the successfull reward session */
	k.updateChallenger(ctx, challenger)

	return &types.MsgRunnerChallengeResponse{}, nil
}

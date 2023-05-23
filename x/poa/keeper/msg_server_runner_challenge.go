package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	param "soarchain/app/params"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
)

const challengerType = "v2n"
const multiplier = int(5)

func (k Keeper) updateChallenger(ctx sdk.Context, challenger types.Challenger) {

	scoreIntChallenger, _ := strconv.Atoi(challenger.Score)
	scoreIntChallenger++

	updatedChallenger := types.Challenger{
		PubKey:       challenger.PubKey,
		Address:      challenger.Address,
		Score:        strconv.Itoa(scoreIntChallenger),
		StakedAmount: challenger.StakedAmount,
		NetEarnings:  challenger.NetEarnings, // TBD
		Type:         challenger.Type,
		IpAddr:       challenger.IpAddr,
	}

	k.SetChallenger(ctx, updatedChallenger)
}

func (k Keeper) coolDownMultiplier(ctx sdk.Context, creator string) uint64 {

	vrfData, _, _ := k.CreateVRF(ctx, creator, multiplier)

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

func (k Keeper) totalEarnings(ctx sdk.Context, netEarning string, rewardMultiplier float64, clientCommunicationMode string) sdk.Coin {
	var totalEarnings sdk.Coin
	var epochRewards sdk.Coin

	/** reward cap check for current epoch */
	targetEpochRewardInt, _ := utility.V2NRewardEmissionPerEpoch(ctx, clientCommunicationMode)
	// if targetEpochErr != nil {
	// 	return totalEarnings, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RunnerChallenge][V2NRewardEmissionPerEpoch] failed. Couldn't calculate reward to emission for each reward"+err.Error())
	// }

	targetEpochReward := sdk.NewCoin(param.BondDenom, sdk.NewIntFromUint64(uint64(targetEpochRewardInt)))
	epochData, _ := k.GetEpochData(ctx)
	// if !found {
	// 	return totalEarnings, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RunnerChallenge][GetEpochData] failed. Couldn't get epoch data with the current context.")
	// }

	epochRewards, _ = sdk.ParseCoinNormalized(epochData.EpochRunner)

	/** check reward cap inside the epoch */
	if epochRewards.IsLT(targetEpochReward) {

		/** Calculate reward earned */
		earnedTokenRewardsFloat, _ := k.V2NRewardCalculator(ctx, rewardMultiplier, clientCommunicationMode)
		// if err != nil {
		// 	return totalEarnings, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RunnerChallenge][V2NRewardCalculator] failed. Couldn't calcualte earned rewards."+err.Error())
		// }

		earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
		earnedCoin := sdk.NewCoin(param.BondDenom, earnedRewardsInt)

		netEarnings, _ := sdk.ParseCoinNormalized(netEarning)
		// if err != nil {
		// 	return totalEarnings, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RunnerChallenge][ParseCoinNormalized] failed. Couldn't parse and normalize a cli input for one coin type, due to invalid or an empty string."+err.Error())
		// }
		totalEarnings = netEarnings.Add(earnedCoin)

		//epochErr := k.UpdateEpochRewards(ctx, clientCommunicationMode, earnedCoin)
		// if epochErr != nil {
		// 	return totalEarnings, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RunnerChallenge][UpdateEpochRewards] failed. Couldn't updat epoch reward."+epochErr.Error())
		// }
	} else {
		netEarnings, _ := sdk.ParseCoinNormalized(netEarning)
		// if err != nil {
		// 	return totalEarnings, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RunnerChallenge][ParseCoinNormalized] failed. Couldn't calculate earned rewards."+err.Error())
		// }
		totalEarnings = netEarnings
	}

	return totalEarnings
}

func (k Keeper) updateRunner(ctx sdk.Context, creator string, runnerAddress string) {

	runner, _ := k.GetChallengeableRunner(ctx, runnerAddress)

	rewardMultiplier, newScore := k.rewardAndScore(runner.Score)

	totalEarnings := k.totalEarnings(ctx, runner.NetEarnings, rewardMultiplier, "runner")

	updatedRunner := types.Runner{
		PubKey:             runner.PubKey,
		Address:            runner.Address,
		Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		StakedAmount:       runner.StakedAmount,
		NetEarnings:        totalEarnings.String(),
		IpAddr:             runner.IpAddr,
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, creator), 10),
	}

	k.SetRunner(ctx, updatedRunner)
}

func (k Keeper) updateClient(ctx sdk.Context, msg *types.MsgRunnerChallenge) {

	// ToDo: Set MOTUS mini rewards
	v2nBxAddrCount := len(msg.ClientPubkeys)
	// if v2nBxAddrCount < 1 {
	// 	return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RunnerChallenge][v2nBxAddrCount] failed. Couldn't find client pubkeys in the tx body.")
	// }

	/** All MOTUS mini devices will be rewarded */
	for i := 0; i < v2nBxAddrCount; i++ {

		v2nBxClient, _ := k.GetClient(ctx, msg.ClientPubkeys[i])
		// if !isFound {
		// 	return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[RunnerChallenge][GetClient] failed. v2n-bx client with index: "+strconv.Itoa(i)+" is not registered in the store.")
		// }

		rewardMultiplier, newScore := k.rewardAndScore(v2nBxClient.Score)

		totalEarnings := k.totalEarnings(ctx, v2nBxClient.NetEarnings, rewardMultiplier, "v2n-bx")

		updatedClient := types.Client{
			Index:              v2nBxClient.Index,
			Address:            v2nBxClient.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        totalEarnings.String(),
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(k.coolDownMultiplier(ctx, msg.Creator), 10),
			Type:               v2nBxClient.Type,
		}

		k.SetClient(ctx, updatedClient)

		k.updateMotusWallet(ctx, v2nBxClient.Address, updatedClient)

	}
}

func (k Keeper) updateMotusWallet(ctx sdk.Context, address string, client types.Client) {

	motusWallet, _ := k.GetMotusWallet(ctx, address)
	// if !isFoundWallet {
	// 	return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RunnerChallenge][GetMotusWallet] failed. Couldn't find Motus wallet.")
	// }
	newMotusWallet := types.MotusWallet{
		Index:  motusWallet.Index,
		Client: &client,
	}
	k.SetMotusWallet(ctx, newMotusWallet)
}

func (k msgServer) RunnerChallenge(goCtx context.Context, msg *types.MsgRunnerChallenge) (*types.MsgRunnerChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, _ := k.GetChallengerByType(ctx, msg.Creator, challengerType)
	// if !found {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RunnerChallenge][GetChallengerByType] failed. Only registered challengers with v2n type can initiate this transaction.")
	// }

	k.updateRunner(ctx, msg.Creator, msg.RunnerAddress)

	k.updateClient(ctx, msg)

	/** Update challenger info after the successfull reward session */
	k.updateChallenger(ctx, challenger)

	return &types.MsgRunnerChallengeResponse{}, nil
}

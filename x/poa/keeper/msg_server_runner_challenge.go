package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	param "soarchain/app/params"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
)

func (k msgServer) RunnerChallenge(goCtx context.Context, msg *types.MsgRunnerChallenge) (*types.MsgRunnerChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, isFound := k.GetChallenger(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only registered challengers can initiate this transaction.")
	}

	// Challenger type must be v2n for this operation
	if challenger.Type != "v2n" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only v2n type challengers can initiate this transaction.")
	}

	// Fetch runner from the store
	runner, isFound := k.GetRunner(ctx, msg.RunnerAddress)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target runner is not registered in the store!")
	}

	// Check runner challengeability
	isChallengeable, point, err := utility.IsChallengeable(ctx, runner.Score, runner.LastTimeChallenged, runner.CoolDownTolerance)
	if err != nil {
		return nil, err
	}
	if !isChallengeable {
		pointString := strconv.FormatFloat(point, 'f', -1, 64)
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Runner is not challengeable at the moment! Point is: "+pointString+" with multiplier: "+runner.CoolDownTolerance)
	}

	// Update runner score
	scoreFloat64, err := strconv.ParseFloat(runner.Score, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
	}
	newScore := utility.CalculateScore(scoreFloat64, true)

	// Update rewardMultiplier
	rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

	// reward cap check for current epoch
	targetEpochRewardInt, targetEpochErr := utility.V2NRewardEmissionPerEpoch(ctx, "runner")
	if targetEpochErr != nil {
		return nil, err
	}
	targetEpochReward := sdk.NewCoin(param.BondDenom, sdk.NewIntFromUint64(uint64(targetEpochRewardInt)))

	epochData, _ := k.GetEpochData(ctx)
	var epochRewards sdk.Coin

	epochRewards, _ = sdk.ParseCoinNormalized(epochData.EpochRunner)

	// check reward cap inside the epoch
	var totalEarnings sdk.Coin
	if epochRewards.IsLT(targetEpochReward) {
		// Calculate reward earned
		earnedTokenRewardsFloat, err := k.V2NRewardCalculator(ctx, rewardMultiplier, "runner")
		if err != nil {
			return nil, err
		}
		earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
		earnedCoin := sdk.NewCoin(param.BondDenom, earnedRewardsInt)

		netEarnings, err := sdk.ParseCoinNormalized(runner.NetEarnings)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		totalEarnings = netEarnings.Add(earnedCoin)

		// update epoch rewards
		epochErr := k.UpdateEpochRewards(ctx, "runner", earnedCoin)
		if epochErr != nil {
			return nil, epochErr
		}
	} else {
		netEarnings, err := sdk.ParseCoinNormalized(runner.NetEarnings)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		totalEarnings = netEarnings
	}

	// Generate random coolDownMultiplier
	multiplier := int(5)

	vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
	if vrfErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "VRF error!")
	}

	generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "vrfData.FinalVrv parse error!")
	}

	var coolDownMultiplier uint64
	if generatedNumber > 0 {
		coolDownMultiplier = generatedNumber
	} else {
		coolDownMultiplier = 1
	}

	updatedRunner := types.Runner{
		PubKey:             runner.PubKey,
		Address:            runner.Address,
		Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		StakedAmount:       runner.StakedAmount,
		NetEarnings:        totalEarnings.String(),
		IpAddr:             runner.IpAddr,
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
		GuardAddress:       runner.GuardAddress,
	}

	k.SetRunner(ctx, updatedRunner)

	// Update runner obj in guard
	guard, isFound := k.GetGuard(ctx, runner.GuardAddress)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Guard not found")
	}
	updateGuard := types.Guard{
		Index:         guard.Index,
		GuardId:       guard.GuardId,
		V2XChallenger: guard.V2XChallenger,
		V2NChallenger: guard.V2NChallenger,
		Runner:        &updatedRunner,
	}

	k.SetGuard(ctx, updateGuard)

	// ToDo: Set MOTUS mini rewards
	v2nBxAddrCount := len(msg.ClientPubkeys)

	if v2nBxAddrCount < 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "No client pubkeys found in the tx body!")
	}

	// All MOTUS mini devices will be rewarded
	for i := 0; i < v2nBxAddrCount; i++ {
		v2nBxClient, isFound := k.GetClient(ctx, msg.ClientPubkeys[i])
		if !isFound {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "v2n-bx client with index: "+strconv.Itoa(i)+" is not registered in the store!")
		}
		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(v2nBxClient.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// reward cap check for current epoch
		targetEpochRewardInt, targetEpochErr := utility.V2NRewardEmissionPerEpoch(ctx, "v2n-bx")
		if targetEpochErr != nil {
			return nil, err
		}
		targetEpochReward := sdk.NewCoin(param.BondDenom, sdk.NewIntFromUint64(uint64(targetEpochRewardInt)))

		epochData, _ := k.GetEpochData(ctx)
		epochRewards, _ := sdk.ParseCoinNormalized(epochData.EpochV2NBX)

		// check reward cap inside the epoch
		var totalEarnings sdk.Coin
		if epochRewards.IsLT(targetEpochReward) {
			// Calculate reward earned
			earnedTokenRewardsFloat, err := k.V2NRewardCalculator(ctx, rewardMultiplier, "v2n-bx")
			if err != nil {
				return nil, err
			}
			earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
			earnedCoin := sdk.NewCoin(param.BondDenom, earnedRewardsInt)

			netEarnings, err := sdk.ParseCoinNormalized(v2nBxClient.NetEarnings)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}
			totalEarnings = netEarnings.Add(earnedCoin)

			// update epoch rewards
			epochErr := k.UpdateEpochRewards(ctx, "v2n-bx", earnedCoin)
			if epochErr != nil {
				return nil, epochErr
			}
		} else {
			netEarnings, err := sdk.ParseCoinNormalized(v2nBxClient.NetEarnings)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}
			totalEarnings = netEarnings
		}

		// Generate random coolDownMultiplier
		multiplier := int(5)

		vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
		if vrfErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "VRF error!")
		}

		generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "vrfData.FinalVrv parse error!")
		}

		var coolDownMultiplier uint64
		if generatedNumber > 0 {
			coolDownMultiplier = generatedNumber
		} else {
			coolDownMultiplier = 1
		}

		updatedClient := types.Client{
			Index:              v2nBxClient.Index,
			Address:            v2nBxClient.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        totalEarnings.String(),
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
		}

		k.SetClient(ctx, updatedClient)

		// Update Motus wallet
		motusWallet, isFoundWallet := k.GetMotusWallet(ctx, v2nBxClient.Address)
		if !isFoundWallet {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Motus client wallet with index: "+strconv.Itoa(i)+" not found!")
		}
		newMotusWallet := types.MotusWallet{
			Index:  motusWallet.Index,
			Client: &updatedClient,
		}
		k.SetMotusWallet(ctx, newMotusWallet)

	}

	// Update challenger info after the successfull reward session
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

	return &types.MsgRunnerChallengeResponse{}, nil
}

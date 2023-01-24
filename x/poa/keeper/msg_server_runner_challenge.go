package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

	// Check tx input of v2n communication mode
	if msg.V2NDeviceType != "v2n-bx" && msg.V2NDeviceType != "runner" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "V2N client communication mode is not supported!")
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

	// Check the challenge result
	result := msg.ChallengeResult

	if result == "reward" {
		// Update runner score
		scoreFloat64, err := strconv.ParseFloat(runner.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// Calculate reward earned
		earnedTokenRewardsFloat, err := k.V2NRewardCalculator(ctx, rewardMultiplier, msg.V2NDeviceType)
		if err != nil {
			return nil, err
		}
		earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
		earnedCoin := sdk.NewCoin("soar", earnedRewardsInt)

		netEarnings, err := sdk.ParseCoinNormalized(runner.NetEarnings)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		totalEarnings := netEarnings.Add(earnedCoin)

		// update epoch rewards
		epochErr := k.UpdateEpochRewards(ctx, msg.V2NDeviceType, earnedCoin)
		if epochErr != nil {
			return nil, epochErr
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
			Index:              runner.Index,
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

	} else if result == "punish" {
		// Update runner score
		scoreFloat64, err := strconv.ParseFloat(runner.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, false)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

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
			Index:              runner.Index,
			Address:            runner.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			StakedAmount:       runner.StakedAmount,
			NetEarnings:        runner.NetEarnings,
			IpAddr:             runner.IpAddr,
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
			GuardAddress:       runner.GuardAddress,
		}

		k.SetRunner(ctx, updatedRunner)

		// Update runner obj in guard
		guard, _ := k.GetGuard(ctx, runner.GuardAddress)
		updateGuard := types.Guard{
			Index:         guard.Index,
			GuardId:       guard.GuardId,
			V2XChallenger: guard.V2XChallenger,
			V2NChallenger: guard.V2NChallenger,
			Runner:        &updatedRunner,
		}

		k.SetGuard(ctx, updateGuard)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenge result")
	}

	// Update challenger info after the successfull reward session
	scoreIntChallenger, _ := strconv.Atoi(challenger.Score)
	scoreIntChallenger++

	updatedChallenger := types.Challenger{
		Index:        challenger.Index,
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

package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
)

func (k msgServer) V2NChallenge(goCtx context.Context, msg *types.MsgV2NChallenge) (*types.MsgV2NChallengeResponse, error) {
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

	// Check the challenge result
	result := msg.RunnerResult

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
		earnedTokenRewards, err := k.V2NRewardCalculator(ctx, rewardMultiplier, "runner")
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		netEarnings, err := strconv.ParseFloat(runner.NetEarnings, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		earnedRewards := netEarnings + earnedTokenRewards

		updatedRunner := types.Runner{
			Index:              runner.Index,
			Address:            runner.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			StakedAmount:       runner.StakedAmount,
			NetEarnings:        strconv.FormatFloat(earnedRewards, 'f', -1, 64),
			IpAddr:             runner.IpAddr,
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetRunner(ctx, updatedRunner)

	} else if result == "punish" {
		// Update runner score
		scoreFloat64, err := strconv.ParseFloat(runner.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, false)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		updatedRunner := types.Runner{
			Index:              runner.Index,
			Address:            runner.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			StakedAmount:       runner.StakedAmount,
			NetEarnings:        runner.NetEarnings,
			IpAddr:             runner.IpAddr,
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetRunner(ctx, updatedRunner)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenge result")
	}

	// V2N-BX
	v2nBxAddrCount := len(msg.V2NBxAddress)
	v2nBxResultCount := len(msg.V2NBxResult)

	if v2nBxAddrCount != v2nBxResultCount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Mismatch between V2N BX address and BX result fields")
	}

	for i := 0; i < v2nBxAddrCount; i++ {
		// Fetch Rx client from the store
		v2nBxClient, isFound := k.GetClient(ctx, msg.V2NBxAddress[i])
		if !isFound {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Bx client is not registered in the store!")
		}
		result := msg.V2NBxResult[i]
		if result == "reward" {
			// Update challengee score
			scoreFloat64, err := strconv.ParseFloat(v2nBxClient.Score, 64)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
			}
			newScore := utility.CalculateScore(scoreFloat64, true)

			// Update rewardMultiplier
			rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

			// Calculate reward earned
			earnedTokenRewards, err := k.V2NRewardCalculator(ctx, rewardMultiplier, "v2n-bx")
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}
			netEarnings, err := strconv.ParseFloat(v2nBxClient.NetEarnings, 64)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}
			earnedRewards := netEarnings + earnedTokenRewards

			updatedClient := types.Client{
				Index:              v2nBxClient.Index,
				Address:            v2nBxClient.Address,
				Registrant:         v2nBxClient.Registrant,
				Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
				RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
				NetEarnings:        strconv.FormatFloat(earnedRewards, 'f', -1, 64),
				LastTimeChallenged: ctx.BlockTime().String(),
			}

			k.SetClient(ctx, updatedClient)
		} else if result == "punish" {
			// Update challengee score
			scoreFloat64, err := strconv.ParseFloat(v2nBxClient.Score, 64)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
			}
			newScore := utility.CalculateScore(scoreFloat64, false)

			// Update rewardMultiplier
			rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

			updatedClient := types.Client{
				Index:              v2nBxClient.Index,
				Address:            v2nBxClient.Address,
				Registrant:         v2nBxClient.Registrant,
				Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
				RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
				NetEarnings:        v2nBxClient.NetEarnings,
				LastTimeChallenged: ctx.BlockTime().String(),
			}

			k.SetClient(ctx, updatedClient)

		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenge result")
		}
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

	return &types.MsgV2NChallengeResponse{}, nil
}

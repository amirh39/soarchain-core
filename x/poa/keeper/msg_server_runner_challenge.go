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

	// Check the challenge result
	runnerAccount, _ := sdk.AccAddressFromBech32(msg.RunnerAddress)

	result := msg.ChallengeResult
	if result == "reward" {
		rewardAmount, _ := sdk.ParseCoinsNormalized("1000000soar")
		//Rewards are issued from the module - soarchain protocol
		k.bankKeeper.MintCoins(ctx, types.ModuleName, rewardAmount)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, runnerAccount, rewardAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
		}

		// Update runner score
		scoreFloat64, err := strconv.ParseFloat(runner.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update runner total rewards
		netEarnings, _ := sdk.ParseCoinsNormalized(runner.NetEarnings)
		rewardAmountCoin, _ := sdk.ParseCoinNormalized("1000000soar")
		netEarnings = netEarnings.Add(rewardAmountCoin)

		updatedRunner := types.Runner{
			Index:              runner.Index,
			Address:            runner.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			StakedAmount:       runner.StakedAmount,
			NetEarnings:        netEarnings.String(),
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

		updatedRunner := types.Runner{
			Index:              runner.Index,
			Address:            runner.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			StakedAmount:       runner.StakedAmount,
			NetEarnings:        runner.NetEarnings,
			IpAddr:             runner.IpAddr,
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetRunner(ctx, updatedRunner)

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

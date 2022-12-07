package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
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

		// increase runner score
		scoreUpdateAmount := 1
		scoreInt, _ := strconv.Atoi(runner.Score)
		scoreInt += scoreUpdateAmount

		// Update runner total rewards
		netEarnings, _ := sdk.ParseCoinsNormalized(runner.NetEarnings)
		rewardAmountCoin, _ := sdk.ParseCoinNormalized("1000000soar")
		netEarnings = netEarnings.Add(rewardAmountCoin)

		updatedRunner := types.Runner{
			Index:        runner.Index,
			Address:      runner.Address,
			Score:        strconv.Itoa(scoreInt),
			StakedAmount: runner.StakedAmount,
			NetEarnings:  netEarnings.String(),
			IpAddr:       runner.IpAddr,
		}

		k.SetRunner(ctx, updatedRunner)

	} else if result == "punish" {
		// decrease runner score
		scoreUpdateAmount := 2
		scoreInt, _ := strconv.Atoi(runner.Score)
		scoreInt -= scoreUpdateAmount

		updatedRunner := types.Runner{
			Index:        runner.Index,
			Address:      runner.Address,
			Score:        strconv.Itoa(scoreInt),
			StakedAmount: runner.StakedAmount,
			NetEarnings:  runner.NetEarnings,
			IpAddr:       runner.IpAddr,
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

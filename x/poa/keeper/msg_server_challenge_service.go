package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//
// Check if sender is a registered validator
// Check the result, reward or punish
// 		 . If reward: mint & send the rewarded coin and increase score
//		 . If punish: decrease score
// Updating the challengee info
// Uptadating challenger info

func (k msgServer) ChallengeService(goCtx context.Context, msg *types.MsgChallengeService) (*types.MsgChallengeServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, isFound := k.GetChallenger(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only registered challengers can initiate this transaction.")
	}

	// Challenger type must be v2x for this operation
	if challenger.Type != "v2x" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only v2x type challengers can initiate this transaction.")
	}

	// Fetch client from the store
	client, isFound := k.GetClient(ctx, msg.ChallengeeAddress)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target client is not registered in the store!")
	}

	// Check the challenge result
	clientAccount, _ := sdk.AccAddressFromBech32(msg.ChallengeeAddress)

	result := msg.ChallengeResult
	if result == "reward" { // reward condition
		rewardAmount, _ := sdk.ParseCoinsNormalized("10000000soar")
		//Rewards are issued from the module - soarchain protocol
		k.bankKeeper.MintCoins(ctx, types.ModuleName, rewardAmount)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, rewardAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
		}

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// Calculate reward earned
		earnedTokenRewards, err := k.MotusReward(ctx, rewardMultiplier)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}

		netEarnings, err := strconv.ParseFloat(client.NetEarnings, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		earnedRewards := netEarnings + earnedTokenRewards

		//
		updatedClient := types.Client{
			Index:              client.Index,
			Address:            client.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        strconv.FormatFloat(earnedRewards, 'f', -1, 64),
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetClient(ctx, updatedClient)

	} else if result == "punish" {
		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, false)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		updatedClient := types.Client{
			Index:              client.Index,
			Address:            client.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        client.NetEarnings,
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetClient(ctx, updatedClient)

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

	return &types.MsgChallengeServiceResponse{}, nil
}

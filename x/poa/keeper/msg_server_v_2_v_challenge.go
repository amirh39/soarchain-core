package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) V2VChallenge(goCtx context.Context, msg *types.MsgV2VChallenge) (*types.MsgV2VChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, isFound := k.GetChallenger(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only registered challengers can initiate this transaction.")
	}

	// Challenger type must be v2x for this operation
	if challenger.Type != "v2x" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only v2x type challengers can initiate this transaction.")
	}

	// Fetch Rx client from the store
	rxClient, isFound := k.GetClient(ctx, msg.RxAddress)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Rx client is not registered in the store!")
	}

	// Check the challenge result
	// clientAccount, _ := sdk.AccAddressFromBech32(msg.RxAddress)
	result := msg.RxResult

	if result == "reward" { // reward condition

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(rxClient.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// Calculate reward earned
		earnedTokenRewards, err := k.V2VRewardCalculator(ctx, rewardMultiplier, "v2v-rx")
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}

		netEarnings, err := strconv.ParseFloat(rxClient.NetEarnings, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
		}
		earnedRewards := netEarnings + earnedTokenRewards

		//
		updatedClient := types.Client{
			Index:              rxClient.Index,
			Address:            rxClient.Address,
			Registrant:         rxClient.Registrant,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        strconv.FormatFloat(earnedRewards, 'f', -1, 64),
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetClient(ctx, updatedClient)

	} else if result == "punish" {
		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(rxClient.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, false)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		updatedClient := types.Client{
			Index:              rxClient.Index,
			Address:            rxClient.Address,
			Registrant:         rxClient.Registrant,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        rxClient.NetEarnings,
			LastTimeChallenged: ctx.BlockTime().String(),
		}

		k.SetClient(ctx, updatedClient)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenge result")
	}

	// BX
	bxAddrCount := len(msg.BxAddress)
	bxResultsCount := len(msg.BxResult)

	if bxAddrCount != bxResultsCount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Mismatch between BX address and BX result fields")
	}

	for i := 0; i < bxAddrCount; i++ {

		// Fetch Rx client from the store
		bxClient, isFound := k.GetClient(ctx, msg.BxAddress[i])
		if !isFound {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Rx client is not registered in the store!")
		}

		result := msg.BxResult[i]
		if result == "reward" {
			// Update challengee score
			scoreFloat64, err := strconv.ParseFloat(msg.BxAddress[i], 64)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
			}
			newScore := utility.CalculateScore(scoreFloat64, true)

			// Update rewardMultiplier
			rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

			// Calculate reward earned
			earnedTokenRewards, err := k.V2VRewardCalculator(ctx, rewardMultiplier, "v2v-bx")
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}

			netEarnings, err := strconv.ParseFloat(bxClient.NetEarnings, 64)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}
			earnedRewards := netEarnings + earnedTokenRewards

			updatedClient := types.Client{
				Index:              bxClient.Index,
				Address:            bxClient.Address,
				Registrant:         bxClient.Registrant,
				Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
				RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
				NetEarnings:        strconv.FormatFloat(earnedRewards, 'f', -1, 64),
				LastTimeChallenged: ctx.BlockTime().String(),
			}

			k.SetClient(ctx, updatedClient)

		} else if result == "punish" {
			// Update challengee score
			scoreFloat64, err := strconv.ParseFloat(bxClient.Score, 64)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
			}
			newScore := utility.CalculateScore(scoreFloat64, false)

			// Update rewardMultiplier
			rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

			updatedClient := types.Client{
				Index:              bxClient.Index,
				Address:            bxClient.Address,
				Registrant:         bxClient.Registrant,
				Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
				RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
				NetEarnings:        bxClient.NetEarnings,
				LastTimeChallenged: ctx.BlockTime().String(),
			}

			k.SetClient(ctx, updatedClient)

		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenge result")
		}
	}

	return &types.MsgV2VChallengeResponse{}, nil
}

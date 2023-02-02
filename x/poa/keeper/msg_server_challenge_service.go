package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
	client, isFound := k.GetClient(ctx, msg.ClientPubkey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target client is not registered in the store!")
	}

	// Check tx input of client communication mode
	if msg.ClientCommunicationMode != "v2v-rx" && msg.ClientCommunicationMode != "v2v-bx" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "V2V client communication mode is not supported!")
	}

	// Check challengeability
	isChallengeable, point, err := utility.IsChallengeable(ctx, client.Score, client.LastTimeChallenged, client.CoolDownTolerance)
	if err != nil {
		return nil, err
	}
	if !isChallengeable {
		pointString := strconv.FormatFloat(point, 'f', -1, 64)
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Client is not challengeable at the moment! Point is: "+pointString+" with multiplier: "+client.CoolDownTolerance)
	}

	// Check the challenge result
	result := msg.ChallengeResult

	if result == "reward" { // reward condition

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot convert to Float64")
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// reward cap check for current epoch
		targetEpochRewardInt, targetEpochErr := utility.V2VRewardEmissionPerEpoch(ctx, msg.ClientCommunicationMode)
		if targetEpochErr != nil {
			return nil, err
		}
		targetEpochReward := sdk.NewCoin("soar", sdk.NewIntFromUint64(uint64(targetEpochRewardInt)))

		epochData, _ := k.GetEpochData(ctx)
		var epochRewards sdk.Coin

		if msg.ClientCommunicationMode == "v2v-rx" {
			epochRewards, _ = sdk.ParseCoinNormalized(epochData.EpochV2VRX)
		} else if msg.ClientCommunicationMode == "v2v-bx" {
			epochRewards, _ = sdk.ParseCoinNormalized(epochData.EpochV2VBX)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "epoch rewards can't be computed because of invalid v2v type!")
		}

		// check reward cap inside the epoch
		var totalEarnings sdk.Coin
		if epochRewards.IsLT(targetEpochReward) {
			// Calculate reward earned
			earnedTokenRewardsFloat, err := k.V2VRewardCalculator(ctx, rewardMultiplier, msg.ClientCommunicationMode)
			if err != nil {
				return nil, err
			}
			earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
			earnedCoin := sdk.NewCoin("soar", earnedRewardsInt)

			netEarnings, err := sdk.ParseCoinNormalized(client.NetEarnings)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot calculate earned rewards!")
			}
			totalEarnings = netEarnings.Add(earnedCoin)

			// update epoch rewards
			epochErr := k.UpdateEpochRewards(ctx, msg.ClientCommunicationMode, earnedCoin)
			if epochErr != nil {
				return nil, epochErr
			}
		} else {
			netEarnings, err := sdk.ParseCoinNormalized(client.NetEarnings)
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
			Index:              client.Index,
			Address:            client.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        totalEarnings.String(),
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
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

		//
		updatedClient := types.Client{
			Index:              client.Index,
			Address:            client.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        client.NetEarnings,
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
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

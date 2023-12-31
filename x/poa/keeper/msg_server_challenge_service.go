package keeper

import (
	"context"
	"log"
	"strconv"

	"github.com/amirh39/soarchain-core/x/poa/constants"
	"github.com/amirh39/soarchain-core/x/poa/types"
	"github.com/amirh39/soarchain-core/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ChallengeService(goCtx context.Context, msg *types.MsgChallengeService) (*types.MsgChallengeServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Challenge Service Transaction Started ##############")

	challenger, isFound := k.GetReputationsByAddress(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[ChallengeService][GetReputationsByAddress] failed. Only registered challengers can initiate this transaction.")
	}

	// Challenger type must be v2x for this operation
	if challenger.Type != constants.V2XChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[ChallengeService][GetChallenger][v2x] failed. Only v2x type challengers can initiate this transaction.")
	}

	if logger != nil {
		logger.Info("V2X challenger is found successfully.", "transaction", "ChallengeService")
	}

	// Fetch client from the store
	client, isFound := k.GetReputation(ctx, msg.ClientPubkey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[ChallengeService][GetClient] failed. Target client is not registered in the store.")
	}

	// Check tx input of client communication mode
	if msg.ClientCommunicationMode != constants.V2VRX && msg.ClientCommunicationMode != constants.V2VBX {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "[ChallengeService][ClientCommunicationMode] failed. V2V client communication mode is not supported.")
	}

	// Check challengeability
	isChallengeable, point, err := utility.IsChallengeable(ctx, client.Score, client.LastTimeChallenged, client.CoolDownTolerance)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "[ChallengeService][IsChallengeable] failed. The target client must be challengeable.")
	}
	if !isChallengeable {
		pointString := strconv.FormatFloat(point, 'f', -1, 64)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][IsChallengeable] failed. Client is not challengeable at the moment. The Point is: "+pointString+" with multiplier: "+client.CoolDownTolerance)
	}

	if logger != nil {
		logger.Info("Checking challengeability for the clinet successfully done.", "transaction", "ChallengeService")
	}

	// Check the challenge result
	result := msg.ChallengeResult

	if result == constants.Reward { // reward condition

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][ParseFloat] failed. Couldn't convert client score to Float64."+err.Error())
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// // reward cap check for current epoch
		// targetEpochRewardInt, targetEpochErr := utility.V2VRewardEmissionPerEpoch(ctx, msg.ClientCommunicationMode)
		// if targetEpochErr != nil {
		// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][V2VRewardEmissionPerEpoch] failed. Couldn't emission reward per epoch.")
		// }
		// targetEpochReward := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetEpochRewardInt)))

		// epochData, found := k.GetEpochData(ctx)
		// if !found {
		// 	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][GetEpochData] failed. Couldn't find epoch data."+err.Error())
		// }

		// var epochRewards sdk.Coin

		// if msg.ClientCommunicationMode == constants.V2VRX {
		// 	epochRewards, err = sdk.ParseCoinNormalized(epochData.EpochV2VRX)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize coin for v2v-rx."+err.Error())
		// 	}
		// } else if msg.ClientCommunicationMode == constants.V2VBX {
		// 	epochRewards, err = sdk.ParseCoinNormalized(epochData.EpochV2VBX)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize coin for v2v-bx."+err.Error())
		// 	}
		// } else {
		// 	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "[ChallengeService] failed. Epoch rewards couldn't be calculated due to invalid v2v type.")
		// }

		// check reward cap inside the epoch
		// var totalEarnings sdk.Coin
		// if epochRewards.IsLT(targetEpochReward) {
		// 	//Calculate reward earned
		// 	earnedTokenRewardsFloat, err := k.V2VRewardCalculator(ctx, rewardMultiplier, msg.ClientCommunicationMode)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ChallengeService][V2VRewardCalculator] failed. Couldn't calculate v2v earned reward.")
		// 	}
		// 	earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
		// 	earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)

		// 	netEarnings, err := sdk.ParseCoinNormalized(client.NetEarnings)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize client new earned coin.")
		// 	}
		// 	totalEarnings = netEarnings.Add(earnedCoin)

		// 	//update epoch rewards
		// 	epochErr := k.UpdateEpochRewards(ctx, msg.ClientCommunicationMode, earnedCoin)
		// 	if epochErr != nil {
		// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][UpdateEpochRewards] failed. Couldn't update epoch rewards.")
		// 	}
		// } else {
		// 	netEarnings, err := sdk.ParseCoinNormalized(client.NetEarnings)
		// 	if err != nil {
		// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize client net earning.")
		// 	}
		// 	totalEarnings = netEarnings
		// }

		// Generate random coolDownMultiplier
		multiplier := int(constants.Multiplier)

		VrfData, err := k.CreateVRF(ctx, msg.Creator, multiplier)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][CreateVRF] failed. Couldn't create a new VRF.")
		}

		generatedNumber, err := strconv.ParseUint(VrfData.FinalVrv, 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][ParseUint] failed. vrfData.FinalVrv parse error.")
		}

		var coolDownMultiplier uint64
		if generatedNumber > 0 {
			coolDownMultiplier = generatedNumber
		} else {
			coolDownMultiplier = 1
		}

		updatedClient := types.Reputation{
			PubKey:             client.PubKey,
			Address:            client.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        "0",
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
		}

		k.SetReputation(ctx, updatedClient)

		if logger != nil {
			logger.Info("Rewarding the client successfully done.", "transaction", "ChallengeService")
		}

	} else if result == constants.Punish {

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][ParseFloat] failed. Cannot convert client score to Float64.")
		}
		newScore := utility.CalculateScore(scoreFloat64, false)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// Generate random coolDownMultiplier
		multiplier := int(5)

		VrfData, err := k.CreateVRF(ctx, msg.Creator, multiplier)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][CreateVRF] failed. Couldn't create a new VRF.")
		}

		generatedNumber, err := strconv.ParseUint(VrfData.FinalVrv, 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[ChallengeService][ParseUint] failed. vrfData.FinalVrv parse error.")
		}

		var coolDownMultiplier uint64
		if generatedNumber > 0 {
			coolDownMultiplier = generatedNumber
		} else {
			coolDownMultiplier = 1
		}

		//
		updatedClient := types.Reputation{
			PubKey:             client.PubKey,
			Address:            client.Address,
			Score:              strconv.FormatFloat(newScore, 'f', -1, 64),
			RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
			NetEarnings:        client.NetEarnings,
			LastTimeChallenged: ctx.BlockTime().String(),
			CoolDownTolerance:  strconv.FormatUint(coolDownMultiplier, 10),
		}

		k.SetReputation(ctx, updatedClient)

		if logger != nil {
			logger.Info("Punishing the client successfully done.", "transaction", "ChallengeService")
		}

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[ChallengeService] failed. Invalid challenge result.")
	}

	// Update challenger info after the successfull reward session
	scoreIntChallenger, _ := strconv.Atoi(challenger.Score)
	scoreIntChallenger++

	updatedChallenger := types.Reputation{
		PubKey:  challenger.PubKey,
		Address: challenger.Address,
		Score:   strconv.Itoa(scoreIntChallenger),
	}

	k.SetReputation(ctx, updatedChallenger)

	if logger != nil {
		logger.Info("Increasing Challenger score successfully done.", "transaction", "ChallengeService")
	}

	log.Println("############## End of Challenge Service Transaction ##############")

	return &types.MsgChallengeServiceResponse{}, nil
}

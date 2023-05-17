package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	params "soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ChallengeService(goCtx context.Context, msg *types.MsgChallengeService) (*types.MsgChallengeServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, isFound := k.GetChallenger(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[ChallengeService][GetChallenger] failed. Only registered challengers can initiate this transaction.")
	}

	// Challenger type must be v2x for this operation
	if challenger.Type != "v2x" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[ChallengeService][GetChallenger][v2x] failed. Only v2x type challengers can initiate this transaction.")
	}

	// Fetch client from the store
	client, isFound := k.GetClient(ctx, msg.ClientPubkey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[ChallengeService][GetClient] failed. Target client is not registered in the store.")
	}

	// Check tx input of client communication mode
	if msg.ClientCommunicationMode != "v2v-rx" && msg.ClientCommunicationMode != "v2v-bx" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "[ChallengeService][ClientCommunicationMode] failed. V2V client communication mode is not supported.")
	}

	// Check challengeability
	isChallengeable, point, err := utility.IsChallengeable(ctx, client.Score, client.LastTimeChallenged, client.CoolDownTolerance)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "[ChallengeService][IsChallengeable] failed. The target client must be challengeable.")
	}
	if !isChallengeable {
		pointString := strconv.FormatFloat(point, 'f', -1, 64)
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][IsChallengeable] failed. Client is not challengeable at the moment. The Point is: "+pointString+" with multiplier: "+client.CoolDownTolerance)
	}

	// Check the challenge result
	result := msg.ChallengeResult

	if result == "reward" { // reward condition

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][ParseFloat] failed. Couldn't convert client score to Float64."+err.Error())
		}
		newScore := utility.CalculateScore(scoreFloat64, true)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// reward cap check for current epoch
		targetEpochRewardInt, targetEpochErr := utility.V2VRewardEmissionPerEpoch(ctx, msg.ClientCommunicationMode)
		if targetEpochErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][V2VRewardEmissionPerEpoch] failed. Couldn't emission reward per epoch."+err.Error())
		}
		targetEpochReward := sdk.NewCoin(params.BondDenom, sdk.NewIntFromUint64(uint64(targetEpochRewardInt)))

		epochData, found := k.GetEpochData(ctx)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][GetEpochData] failed. Couldn't find epoch data."+err.Error())
		}

		var epochRewards sdk.Coin

		if msg.ClientCommunicationMode == "v2v-rx" {
			epochRewards, err = sdk.ParseCoinNormalized(epochData.EpochV2VRX)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize coin for v2v-rx."+err.Error())
			}
		} else if msg.ClientCommunicationMode == "v2v-bx" {
			epochRewards, err = sdk.ParseCoinNormalized(epochData.EpochV2VBX)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize coin for v2v-bx."+err.Error())
			}
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "[ChallengeService] failed. Epoch rewards couldn't be calculated due to invalid v2v type.")
		}

		// check reward cap inside the epoch
		var totalEarnings sdk.Coin
		if epochRewards.IsLT(targetEpochReward) {
			// Calculate reward earned
			earnedTokenRewardsFloat, err := k.V2VRewardCalculator(ctx, rewardMultiplier, msg.ClientCommunicationMode)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[ChallengeService][V2VRewardCalculator] failed. Couldn't calculate v2v earned reward."+err.Error())
			}
			earnedRewardsInt := sdk.NewIntFromUint64((uint64(earnedTokenRewardsFloat)))
			earnedCoin := sdk.NewCoin(params.BondDenom, earnedRewardsInt)

			netEarnings, err := sdk.ParseCoinNormalized(client.NetEarnings)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize client new earned coin."+err.Error())
			}
			totalEarnings = netEarnings.Add(earnedCoin)

			// update epoch rewards
			epochErr := k.UpdateEpochRewards(ctx, msg.ClientCommunicationMode, earnedCoin)
			if epochErr != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][UpdateEpochRewards] failed. Couldn't update epoch rewards."+epochErr.Error())
			}
		} else {
			netEarnings, err := sdk.ParseCoinNormalized(client.NetEarnings)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][ParseCoinNormalized] failed. Couldn't parse and normalize client net earning."+err.Error())
			}
			totalEarnings = netEarnings
		}

		// Generate random coolDownMultiplier
		multiplier := int(5)

		vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
		if vrfErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][CreateVRF] failed. Couldn't create a new VRF."+vrfErr.Error())
		}

		generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][ParseUint] failed. vrfData.FinalVrv parse error."+err.Error())
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
			Type:               client.Type,
		}

		k.SetClient(ctx, updatedClient)

		// Update Motus wallet
		motusWallet, isFoundWallet := k.GetMotusWallet(ctx, client.Address)
		if !isFoundWallet {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][GetMotusWallet] failed. Couldn't find a wallet for Motus client.")
		}
		newMotusWallet := types.MotusWallet{
			Index:  motusWallet.Index,
			Client: &updatedClient,
		}
		k.SetMotusWallet(ctx, newMotusWallet)

	} else if result == "punish" {

		// Update challengee score
		scoreFloat64, err := strconv.ParseFloat(client.Score, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][ParseFloat] failed. Cannot convert client score to Float64."+err.Error())
		}
		newScore := utility.CalculateScore(scoreFloat64, false)

		// Update rewardMultiplier
		rewardMultiplier := utility.CalculateRewardMultiplier(newScore)

		// Generate random coolDownMultiplier
		multiplier := int(5)

		vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
		if vrfErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][CreateVRF] failed. Couldn't create a new VRF."+vrfErr.Error())
		}

		generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[ChallengeService][ParseUint] failed. vrfData.FinalVrv parse error."+err.Error())
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
			Type:               client.Type,
		}

		k.SetClient(ctx, updatedClient)

		// Update Motus wallet
		motusWallet, isFoundWallet := k.GetMotusWallet(ctx, client.Address)
		if !isFoundWallet {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[ChallengeService][GetMotusWallet] failed. Motus client wallet not found.")
		}
		newMotusWallet := types.MotusWallet{
			Index:  motusWallet.Index,
			Client: &updatedClient,
		}
		k.SetMotusWallet(ctx, newMotusWallet)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[ChallengeService] failed. Invalid challenge result.")
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

	return &types.MsgChallengeServiceResponse{}, nil
}

package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Tasks:
// 1. Check if sender is a registered validator
// 2. Check the result, reward or punish
// 		 . If reward: mint & send the rewarded coin and increase score
//		 . If punish: decrease score
// 3. updating the challengee info
// 4. uptadating challenger info

func (k msgServer) ChallengeService(goCtx context.Context, msg *types.MsgChallengeService) (*types.MsgChallengeServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, isChallenger := k.GetChallenger(ctx, msg.Creator)
	if !isChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only registered challengers can initiate this transaction.")
	}

	// Try to fetch client from the store
	client, isFound := k.GetClient(ctx, msg.ChallengeeAddress)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target client is not registered in the store!")
	}

	// Check the challenge result
	clientAccount, _ := sdk.AccAddressFromBech32(msg.ChallengeeAddress)

	result := msg.ChallengeResult
	if result == "reward" { // reward condition
		rewardAmount, _ := sdk.ParseCoinsNormalized("10soar")
		//Rewards are issued from the module - soarchain protocol
		k.bankKeeper.MintCoins(ctx, types.ModuleName, rewardAmount)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, rewardAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Cannot send coins")
		}

		// increase challengee score
		scoreUpdateAmount := 5
		scoreInt, _ := strconv.Atoi(client.Score)
		scoreInt += scoreUpdateAmount

		// update challengee total rewards
		netEarnings, _ := sdk.ParseCoinsNormalized(client.NetEarnings)
		rewardAmountCoin, _ := sdk.ParseCoinNormalized("10soar")
		netEarnings = netEarnings.Add(rewardAmountCoin)

		updatedClient := types.Client{
			Index:       client.Index,
			Address:     client.Address,
			UniqueId:    client.UniqueId,
			Score:       strconv.Itoa(scoreInt),
			NetEarnings: netEarnings.String(),
		}

		k.SetClient(ctx, updatedClient)

	} else if result == "punish" {
		// decrease challengee score
		scoreUpdateAmount := 5
		scoreInt, _ := strconv.Atoi(client.Score)
		scoreInt -= scoreUpdateAmount

		// ToDo: implement financial punishment

		updatedClient := types.Client{
			Index:       client.Index,
			Address:     client.Address,
			UniqueId:    client.UniqueId,
			Score:       strconv.Itoa(scoreInt),
			NetEarnings: client.NetEarnings,
		}

		k.SetClient(ctx, updatedClient)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenge result")
	}

	return &types.MsgChallengeServiceResponse{}, nil
}

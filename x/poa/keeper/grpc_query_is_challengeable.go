package keeper

import (
	"context"

	"strconv"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) IsChallengeable(goCtx context.Context, req *types.QueryIsChallengeableRequest) (*types.QueryIsChallengeableResponse, error) {

	if req == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[IsChallengeable] failed. Invalid request: %T.", req)
	}

	if !utility.ValidString(req.ClientAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[IsChallengeable][ValidString] failed. Couldn't find a valid client address from the request. got: [ %T ], Make sure client address is not empty OR invalid.", req.ClientAddr)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reputation, isFound := k.GetReputationsByAddress(ctx, req.ClientAddr)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[IsChallengeable][GetClient] failed. Target client is not registered in the store.")
	}

	// Check challengeability
	isChallengeable, point, err := utility.IsChallengeable(ctx, reputation.Score, reputation.LastTimeChallenged, reputation.CoolDownTolerance)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "[IsChallengeable][utility.IsChallengeable] failed. Couldn't calculate challengeability from the client entity. got [ %T ], Make sure the client entity includes valid Score, LastTimeChallenged and CoolDownTolerance. ", reputation)
	}

	isChallengeableStr := strconv.FormatBool(isChallengeable)
	pointString := strconv.FormatFloat(point, 'f', -1, 64)

	return &types.QueryIsChallengeableResponse{ResultBool: isChallengeableStr, ChallengeabilityScore: pointString}, nil
}

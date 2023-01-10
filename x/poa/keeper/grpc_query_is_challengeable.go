package keeper

import (
	"context"

	"strconv"

	"soarchain/x/poa/types"

	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

func (k Keeper) IsChallengeable(goCtx context.Context, req *types.QueryIsChallengeableRequest) (*types.QueryIsChallengeableResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	client, isFound := k.GetClient(ctx, req.ClientAddr)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Target client is not registered in the store!")
	}

	// Check challengeability
	isChallengeable, point, err := utility.IsChallengeable(ctx, client.Score, client.LastTimeChallenged, client.CoolDownTolerance)
	if err != nil {
		return nil, err
	}

	isChallengeableStr := strconv.FormatBool(isChallengeable)

	pointString := strconv.FormatFloat(point, 'f', -1, 64)

	return &types.QueryIsChallengeableResponse{ResultBool: isChallengeableStr, ChallengeabilityScore: pointString}, nil

}

package keeper

import (
	"context"

	"github.com/amirh39/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyRandomNumber(goCtx context.Context, req *types.QueryVerifyRandomNumberRequest) (*types.QueryVerifyRandomNumberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "[VerifyRandomNumber] failed. Invalid request.")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	isVerified, err := k.VerifyGeneratedNumber(ctx, req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[VerifyRandomNumber][VerifyGeneratedNumber] failed. Invalid generated number with VRF data.")
	}

	return &types.QueryVerifyRandomNumberResponse{Result: isVerified}, nil
}

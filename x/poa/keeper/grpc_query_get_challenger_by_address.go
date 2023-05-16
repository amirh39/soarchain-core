package keeper

import (
	"context"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k Keeper) GetChallengerByAddress(goCtx context.Context, req *types.QueryGetChallengerByAddressRequest) (*types.QueryGetChallengerByAddressResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[GetChallengerByAddress] failed. Invalid request: %T.", req)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, found := k.GetChallenger(ctx, req.Address)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "[GetChallengerByAddress][GetChallenger] failed. Couldn't find a challenger by the this address: [ %T ], Make sure address is not empty OR invalid.", req.Address)
	}

	return &types.QueryGetChallengerByAddressResponse{Challenger: &challenger}, nil
}

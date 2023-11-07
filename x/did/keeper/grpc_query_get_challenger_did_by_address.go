package keeper

import (
	"context"

	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetChallengerDidByAddress(goCtx context.Context, req *types.QueryGetChallengerDidByAddressRequest) (*types.QueryGetChallengerDidByAddressResponse, error) {
	if req == nil || req.Address == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[GetChallengerDidByAddress] failed. Invalid request: %T.", req)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenger, found := k.GetChallengerDid(ctx, req.Address)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "[GetChallengerDidByAddress][GetChallengerDid] failed. Couldn't find a challenger by the this address: [ %s ], Make sure address is not empty OR invalid.", req.Address)
	}

	return &types.QueryGetChallengerDidByAddressResponse{ChallengerDid: &challenger}, nil
}

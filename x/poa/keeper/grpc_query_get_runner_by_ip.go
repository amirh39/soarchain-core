package keeper

import (
	"context"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetRunnerByIp(goCtx context.Context, req *types.QueryGetRunnerByIpRequest) (*types.QueryGetRunnerByIpResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[GetRunnerByIp] failed. Invalid request: %T.", req)
	}

	if !utility.ValidString(req.IpAddress) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[GetRunnerByIp][ValidString] failed. Couldn't find a valid Ip Address from the request. got: [ %T ], Make sure IP Address is not empty OR invalid.", req.IpAddress)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	runners := k.GetAllRunner(ctx)
	if len(runners) == 0 || runners == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GetRunnerByIp][GetAllRunner] failed. Couldn't find any runner.")
	}

	targetRunner := types.Runner{}

	for _, runner := range runners {
		if req.IpAddress == runner.IpAddress {
			targetRunner = runner
			break
		}
	}

	if !utility.ValidString(targetRunner.IpAddress) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[GetRunnerByIp][ValidString] failed. Couldn't find a valid Ip Address from the target runner. got: [ %T ], Make sure IP Address is not empty OR invalid.", targetRunner.IpAddress)
	}

	return &types.QueryGetRunnerByIpResponse{Runner: &targetRunner}, nil
}

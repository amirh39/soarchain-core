package keeper

import (
	"context"

	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetRunnerDidByIp(goCtx context.Context, req *types.QueryGetRunnerDidByIpRequest) (*types.QueryGetRunnerDidByIpResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[GetRunnerByIp] failed. Invalid request: %T.", req)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !ValidString(req.IpAddress) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[GetRunnerDidByIp][ValidString] failed. Couldn't find a valid Ip Address from the request. got: [ %T ], Make sure IP Address is not empty OR invalid.", req.IpAddress)
	}

	runners := k.GetAllRunnerDid(ctx)
	if len(runners) == 0 || runners == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GetRunnerDidByIp][GetAllRunnerDid] failed. Couldn't find any runner did.")
	}

	targetRunner := types.RunnerDid{}

	for _, runner := range runners {
		if req.IpAddress == runner.IpAddress {
			targetRunner = runner
			break
		}
	}

	return &types.QueryGetRunnerDidByIpResponse{Runner: &targetRunner}, nil
}

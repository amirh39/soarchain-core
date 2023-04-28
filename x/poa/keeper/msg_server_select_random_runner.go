package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SelectRandomRunner(goCtx context.Context, msg *types.MsgSelectRandomRunner) (*types.MsgSelectRandomRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allRunners := k.GetAllRunner(ctx)
	if allRunners != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[SelectRandomRunner][GetAllRunner] failed. Couldn't get all runners list.")
	}

	multiplier := int(len(allRunners))

	vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
	if vrfErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[SelectRandomRunner][CreateVRF] failed. VRF error!")
	}

	generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[SelectRandomRunner][ParseUint] failed. vrfData.FinalVrv parse error!"+err.Error())
	}
	var selectedRunner types.Runner
	runenrs := k.GetAllRunner(ctx)
	if allRunners != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[SelectRandomRunner][GetAllRunner] failed. Couldn't get all runners list.")
	}

	for i := 0; i < len(runenrs); i++ {
		if i == int(generatedNumber) {
			selectedRunner = runenrs[i]
		}
	}

	// record selected challenger for future referenece
	vrf, isFound := k.GetVrfData(ctx, vrfData.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[SelectRandomRunner][GetVrfData] failed. Vrf data couldn't found.")
	}
	updateVrf := types.VrfData{
		Index:              vrf.Index,
		Creator:            vrf.Creator,
		Vrv:                vrf.Vrv,
		Multiplier:         vrf.Multiplier,
		Proof:              vrf.Proof,
		Pubkey:             vrf.Pubkey,
		Message:            vrf.Message,
		ParsedVrv:          vrf.ParsedVrv,
		FloatVrv:           vrf.FloatVrv,
		FinalVrv:           vrf.FinalVrv,
		FinalVrvFloat:      vrf.FinalVrvFloat,
		SelectedChallenger: nil,
		SelectedRunner:     &selectedRunner,
	}
	k.SetVrfData(ctx, updateVrf)

	return &types.MsgSelectRandomRunnerResponse{RandomRunner: &selectedRunner}, nil
}

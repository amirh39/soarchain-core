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
	multiplier := int(len(allRunners))

	vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
	if vrfErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "VRF error!")
	}

	generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "vrfData.FinalVrv parse error!")
	}
	var selectedRunner types.Runner
	runenrs := k.GetAllRunner(ctx)
	for i := 0; i < len(runenrs); i++ {
		if i == int(generatedNumber) {
			selectedRunner = runenrs[i]
		}
	}

	// record selected challenger for future referenece
	vrf, isFound := k.GetVrfData(ctx, vrfData.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Vrf data can't be found!")
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

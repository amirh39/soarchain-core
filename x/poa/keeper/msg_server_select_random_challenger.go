package keeper

import (
	"context"
	"strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SelectRandomChallenger(goCtx context.Context, msg *types.MsgSelectRandomChallenger) (*types.MsgSelectRandomChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allChallengers := k.GetAllChallenger(ctx)
	multiplier := int(len(allChallengers))

	vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
	if vrfErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "VRF error!")
	}

	generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "vrfData.FinalVrv parse error!")
	}
	var selectedChallenger types.Challenger
	challengers := k.GetAllChallenger(ctx)
	for i := 0; i < len(challengers); i++ {
		if i == int(generatedNumber) {
			selectedChallenger = challengers[i]
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
		SelectedChallenger: &selectedChallenger,
		SelectedRunner:     nil,
	}
	k.SetVrfData(ctx, updateVrf)

	return &types.MsgSelectRandomChallengerResponse{RandomChallenger: &selectedChallenger}, nil
}

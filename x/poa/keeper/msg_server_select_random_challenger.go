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

	if msg.Creator == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[SelectRandomChallenger] failed. Couldn't find a valid msg.Creator. got [ %T ]", msg.Creator)
	}

	allChallengers := k.GetAllChallenger(ctx)
	if allChallengers == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[SelectRandomChallenger][GetAllChallenger] failed. Couldn't find any challenger.")
	}
	multiplier := int(len(allChallengers))

	vrfData, _, vrfErr := k.CreateVRF(ctx, msg.Creator, multiplier)
	if vrfErr != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPanic, "[SelectRandomChallenger][CreateVRF] failed. VRF couldn't create. Error [ %T ]", vrfErr)
	}

	generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPanic, "[SelectRandomChallenger][ParseUint] failed. Couldn't parse VRF data. Error: [ %T ]", err)
	}
	var selectedChallenger types.Challenger
	challengers := k.GetAllChallenger(ctx)
	if allChallengers == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[SelectRandomChallenger][GetAllChallenger] failed. Couldn't find any challenger.")
	}
	for i := 0; i < len(challengers); i++ {
		if i == int(generatedNumber) {
			selectedChallenger = challengers[i]
		}
	}

	// record selected challenger for future referenece
	vrf, isFound := k.GetVrfData(ctx, vrfData.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[SelectRandomChallenger][GetVrfData] failed. Couldn't get VRF data.")
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

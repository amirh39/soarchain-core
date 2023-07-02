package keeper

import (
	"context"
	"log"
	"strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SelectRandomRunner(goCtx context.Context, msg *types.MsgSelectRandomRunner) (*types.MsgSelectRandomRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Select Random Runner Transaction Started ##############")

	if msg.Creator == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[SelectRandomRunner] failed. Couldn't find a valid msg.Creator. got [ %T ]", msg.Creator)
	}

	runenrs := k.GetAllRunner(ctx)
	if runenrs == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[SelectRandomRunner][GetAllRunner] failed. Couldn't find any runner.")
	}

	factor := int(len(runenrs))

	VrfData, err := k.CreateVRF(ctx, msg.Creator, factor)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[SelectRandomRunner][CreateVRF] failed. Error: [ %T ]", err)
	}

	generatedNumber, err := strconv.ParseUint(VrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[SelectRandomRunner][ParseUint] failed. vrfData.FinalVrv parse error: [ %T ]", err)
	}

	var selectedRunner types.Runner
	for i := 0; i < len(runenrs); i++ {
		if i == int(generatedNumber) {
			selectedRunner = runenrs[i]
		}
	}

	// record selected challenger for future referenece
	vrf, isFound := k.GetVrfData(ctx, VrfData.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[SelectRandomRunner][GetVrfData] failed. Vrf data couldn't found.")
	}

	updateVrf := types.VrfData{
		Index:         vrf.Index,
		Creator:       vrf.Creator,
		Vrv:           vrf.Vrv,
		Multiplier:    vrf.Multiplier,
		Proof:         vrf.Proof,
		Pubkey:        vrf.Pubkey,
		Message:       vrf.Message,
		ParsedVrv:     vrf.ParsedVrv,
		FloatVrv:      vrf.FloatVrv,
		FinalVrv:      vrf.FinalVrv,
		FinalVrvFloat: vrf.FinalVrvFloat,
	}
	k.SetVrfData(ctx, updateVrf)

	if logger != nil {
		logger.Info("Updating Vrf data successfully done.", "transaction", "SelectRandomRunner")
	}

	log.Println("############## End of Select Random Runner Transaction ##############")

	return &types.MsgSelectRandomRunnerResponse{RandomRunner: &selectedRunner}, nil
}

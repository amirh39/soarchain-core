package keeper

import (
	"context"
	"log"
	"strconv"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SelectRandomChallenger(goCtx context.Context, msg *types.MsgSelectRandomChallenger) (*types.MsgSelectRandomChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Select Random Challenger Transaction Started ##############")

	if msg.Creator == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[SelectRandomChallenger] failed. Couldn't find a valid msg.Creator. got [ %T ]", msg.Creator)
	}

	allChallengerReputations := k.GetAllChallenger(ctx)
	if allChallengerReputations == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[SelectRandomChallenger][GetAllChallenger] failed. Couldn't find any challenger.")
	}
	factor := int(len(allChallengerReputations))

	vrfData, err := k.CreateVRF(ctx, msg.Creator, factor)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[SelectRandomChallenger][CreateVRF] failed. VRF couldn't create. Error [ %T ]", err)
	}

	generatedNumber, err := strconv.ParseUint(vrfData.FinalVrv, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[SelectRandomChallenger][ParseUint] failed. Couldn't parse VRF data. Error: [ %T ]", err)
	}
	var selectedChallenger types.Reputation
	for i := 0; i < len(allChallengerReputations); i++ {
		if i == int(generatedNumber) {
			selectedChallenger = allChallengerReputations[i]
		}
	}

	// record selected challenger for future referenece
	vrf, isFound := k.GetVrfData(ctx, vrfData.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[SelectRandomChallenger][GetVrfData] failed. Couldn't get VRF data.")
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
		logger.Info("Updating Vrf data successfully done.", "transaction", "SelectRandomChallenger")
	}

	log.Println("############## End of Select Random Challenger Transaction ##############")

	return &types.MsgSelectRandomChallengerResponse{RandomChallengerReputation: &selectedChallenger}, nil
}

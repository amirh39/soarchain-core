package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

func (k msgServer) GenDpr(goCtx context.Context, msg *types.MsgGenDpr) (*types.MsgGenDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	epochDtata, found := k.epochKeeper.GetEpochData(ctx)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][GetEpochData] failed. Couldn't find epoch data.")
	}

	result := k.VerifyDprInputs(msg, epochDtata.TotalEpochs)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][VerifyDprInputs] failed. Make sure you are using valid inputs for creating Dpr object.")
	}

	if logger != nil {
		logger.Info("Validationg DPR is successfully Done.", "transaction", "GenDpr")
	}

	// pinNumbers := utility.CalculatePinNumber(msg)

	// eligible := k.Keeper.didKeeper.GetEligibleDids(ctx, pinNumbers)
	// if !eligible {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][GetEligibleDids] failed. There is no eligible client to serve this DPR.")
	// }

	//Save dpr into storage
	newDpr := types.Dpr{
		Id:            uuid.New().String(),
		Creator:       msg.Creator,
		SupportedPIDs: msg.SupportedPIDs,
		IsActive:      false,
		Vin:           []string{},
		ClientPubkeys: []string{},
		Duration:      msg.Duration,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Generating DPR is successfully Done.", "transaction", "GenDpr")
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgGenDprResponse{}, nil
}

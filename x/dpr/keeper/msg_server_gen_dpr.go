package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"
	"soarchain/x/dpr/utility"

	"github.com/pborman/uuid"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenDpr(goCtx context.Context, msg *types.MsgGenDpr) (*types.MsgGenDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a dpr Transaction is Started ##############")

	epochDtata, found := k.epochKeeper.GetEpochData(ctx)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][GetEpochData] failed. Couldn't find epoch data.")
	}

	if logger != nil {
		logger.Info("epochDtata.TotalEpochs--===>", "transaction", "GenDpr", "epochDtata.TotalEpochs", epochDtata.TotalEpochs)
	}

	result := k.VerifyDPRInputs(msg, epochDtata.TotalEpochs)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][VerifyDPRInputs] failed. Make sure you are using valid properties for creating Dpr object.")
	}

	if logger != nil {
		logger.Info("Verifying a received DPR sucessfully done.", "transaction", "GenDpr")
	}

	found = k.didKeeper.FindEligibleDid(ctx, utility.CalculatePinNumber(msg))
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][FindEligibleDid] failed. There is no eligible client to serve this DPR.")
	}

	if logger != nil {
		logger.Info("Finding eligible client for serving DPR sucessfully done.", "transaction", "GenDpr", "found", found)
	}

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:                            uuid.NewRandom().String(),
		Creator:                       msg.Creator,
		PidSupportedOneToTwnety:       msg.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: msg.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  msg.PidSupportedForthyOneToSixty,
		IsActive:                      true,
		Vin:                           []string{},
		ClientPubkeys:                 []string{},
		LengthOfDpr:                   msg.LengthOfDpr,
	}

	k.SetDpr(ctx, newDpr)

	xx, _ := k.GetDpr(ctx, newDpr.Id)

	if logger != nil {
		logger.Info("saved dpr id fetched.", "transaction", "GenDpr", "xx", xx)
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgGenDprResponse{}, nil
}

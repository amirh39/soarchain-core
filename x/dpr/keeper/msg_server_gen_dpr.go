package keeper

import (
	"context"
	"fmt"
	"log"

	"soarchain/x/dpr/types"
	"soarchain/x/dpr/utility"

	"github.com/pborman/uuid"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) validateProperties(msg *types.MsgGenDpr) bool {

	result := true

	if msg.Creator == "" {
		result = false
	}

	fmt.Print("22222222222222", result)

	if msg.Vin == nil {
		result = false
	}

	fmt.Print("333333333333", result)

	if !msg.PidSupported_1To_20 && !msg.PidSupported_21To_40 && !msg.PidSupported_41To_60 {
		result = false
	}

	fmt.Print("44444444444444", result)

	return result
}

func (k msgServer) GenDpr(goCtx context.Context, msg *types.MsgGenDpr) (*types.MsgGenDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a dpr Transaction Started ##############")

	result := k.validateProperties(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][validateProperties] failed. Make sure you are using valid properties for creating Dpr object.")
	}

	isActive, err := utility.CalculateDprValidity(msg.LengthOfDpr)
	if !isActive {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][CalculateDprValidity] failed. Dpr is not active.")
	}
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][CalculateDprValidity] failed. Couldn't calculate the time of activation for a Dpr.")
	}

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "GenDpr")
	}

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:                   uuid.NewRandom().String(),
		Creator:              msg.Creator,
		PidSupported_1To_20:  msg.PidSupported_1To_20,
		PidSupported_21To_40: msg.PidSupported_21To_40,
		PidSupported_41To_60: msg.PidSupported_41To_60,
		IsActive:             true,
		Vin:                  msg.Vin,
		ClientPubkeys:        []string{},
	}

	k.SetDpr(ctx, newDpr)

	xx := k.GetAllDpr(ctx)

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "GenDpr", "dpr-objects", xx)
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgGenDprResponse{}, nil
}

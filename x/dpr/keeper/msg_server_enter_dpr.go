package keeper

import (
	"context"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnterDpr(goCtx context.Context, msg *types.MsgEnterDpr) (*types.MsgEnterDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a dpr Transaction Started ##############")

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "EnterDpr")
	}

	dpr, found := k.GetDpr(ctx, msg.DprId)

	did, found := k.didKeeper.GetDidDocumentByPubkey(ctx, msg.PubKey)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[EnterDpr][GetPidSupportedDidDocument] failed. There is no eligible client to serve this DPR.")
	}

	clientPubKey := []string{}
	clientPubKey = append(clientPubKey, msg.PubKey)

	vins := []string{}
	vins = append(vins, did.Document.Vin)

	// Save dpr into storage
	newDpr := types.Dpr{
		Id:                            dpr.Id,
		Creator:                       dpr.Creator,
		PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
		PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
		PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
		IsActive:                      true,
		Vin:                           vins,
		ClientPubkeys:                 clientPubKey,
		LengthOfDpr:                   dpr.LengthOfDpr,
	}

	k.SetDpr(ctx, newDpr)

	// xx := k.GetAllDpr(ctx)

	if logger != nil {
		logger.Info("Dpr is vali and active", "transaction", "EnterDpr", "dpr-objects")
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgEnterDprResponse{}, nil
}

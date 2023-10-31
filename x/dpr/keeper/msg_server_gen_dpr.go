package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenDpr(goCtx context.Context, msg *types.MsgGenDpr) (*types.MsgGenDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	epochData, found := k.epochKeeper.GetEpochData(ctx)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][GetEpochData] failed. Couldn't find epoch data.")
	}

	result := k.VerifyDprInputs(msg, epochData.TotalEpochs)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDpr][VerifyDprInputs] failed. Make sure you are using valid inputs for creating Dpr object.")
	}

	if logger != nil {
		logger.Info("Validating DPR is successfully Done.", "transaction", "GenDpr")
	}

	// Coin denomination check already done
	budget, err := sdk.ParseCoinsNormalized(msg.DprBudget)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[GenDpr][ParseCoinsNormalized] failed. Couldn't parse budget.")
	}

	dprOwner, _ := sdk.AccAddressFromBech32(msg.Creator)

	errTransfer := k.bankKeeper.SendCoinsFromAccountToModule(ctx, dprOwner, types.ModuleName, budget)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[GenDpr][SendCoinsFromAccountToModule] failed. Couldn't send coins.")
	}

	timestampStr := ctx.BlockTime().String()
	hashPayload := msg.Creator + "|" + timestampStr
	hash := sha256.Sum256([]byte(hashPayload))
	dprID := hex.EncodeToString(hash[:])

	//Save dpr into storage
	newDpr := types.Dpr{
		Id:             dprID,
		Creator:        msg.Creator,
		SupportedPIDs:  msg.SupportedPIDs,
		IsActive:       false,
		ClientPubkeys:  []string{},
		Duration:       msg.Duration,
		DprEndTime:     "",
		DprStartEpoch:  0,
		DprBudget:      budget.String(),
		MaxClientCount: msg.MaxClientCount,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Generating DPR is successfully Done.", "transaction", "GenDpr")
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgGenDprResponse{}, nil
}

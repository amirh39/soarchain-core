package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/amirh39/soarchain-core/app/params"
	"github.com/amirh39/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenDpr(goCtx context.Context, msg *types.MsgGenDpr) (*types.MsgGenDprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	result := k.VerifyDprInputs(msg)
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

	// Calculate 1% of the budget amount for the specific denomination
	onePercentAmt := budget.AmountOf(params.BondDenom).QuoRaw(100)

	// Create a coin with 1% of the budget
	onePercentCoin := sdk.NewCoin(params.BondDenom, onePercentAmt)

	// Subtract 1% from the original budget
	remainingBudget := budget.Sub(sdk.NewCoins(onePercentCoin))
	if !remainingBudget.IsValid() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Subtracting %1 from the budget results in invalid coins.")
	}

	// TODO: find an efficient way to distribute rewards to runners

	dprOwner, _ := sdk.AccAddressFromBech32(msg.Creator)
	//log.Println(k.bankKeeper.GetBalance(ctx, dprOwner, params.BondDenom))
	errTransfer := k.bankKeeper.SendCoinsFromAccountToModule(ctx, dprOwner, types.ModuleName, budget)
	if errTransfer != nil {
		return nil, errTransfer
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
		Status:         0,
		Duration:       msg.Duration,
		DprEndTime:     "",
		DprStartEpoch:  0,
		DprBudget:      budget.String(),
		MaxClientCount: msg.MaxClientCount,
		Name:           msg.Name,
	}
	k.SetDpr(ctx, newDpr)

	if logger != nil {
		logger.Info("Generating DPR is successfully Done.", "transaction", "GenDpr")
	}

	log.Println("############## End of Generating dpr Transaction ##############")

	return &types.MsgGenDprResponse{}, nil
}

package keeper

import (
	"context"
	"log"

	"github.com/soar-robotics/soarchain-core/app/params"
	"github.com/soar-robotics/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimDprRewards(goCtx context.Context, msg *types.MsgClaimDprRewards) (*types.MsgClaimDprRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	log.Println("############## Claim DPR Rewards Transaction Started ##############")

	did, isFound := k.didKeeper.GetClientDid(ctx, msg.Sender)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[ClaimDprRewards][GetClientDid] failed. Creator is not a valid address. Only motus owners can send this transaction.")
	}

	earnings, err := k.DistributeRewards(ctx, did)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "[ClaimDprRewards][calculateDprEarnings] failed.")
	}

	// Update the claimed amount with the new earnings
	reputation, isFound := k.poaKeeper.GetReputation(ctx, did.PubKey)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[ClaimDprRewards][GetReputation] failed.")
	}
	log.Printf("Current reputation: %+v\n", reputation)

	// Parse the existing DPR earnings from the reputation
	existingDprEarnings, err := sdk.ParseCoinsNormalized(reputation.DprEarnings)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Failed to parse existing DPR earnings")
	}
	log.Printf("Existing DPR Earnings: %+v\n", existingDprEarnings)

	// Ensure earnings is of type sdk.DecCoins and get the integer amount
	earningsInt := earnings.AmountOf(params.BondDenom).TruncateInt()

	// Calculate the new DPR earnings by adding the earnings from the existing DPR earnings
	newDprEarnings := existingDprEarnings.Add(sdk.NewCoin(params.BondDenom, earningsInt))

	earningsCurrent := sdk.NewCoins(sdk.NewCoin(params.BondDenom, earningsInt))

	log.Printf("New DPR Earnings after addition: %+v\n", newDprEarnings)

	// Update the reputation with the new DPR earnings
	reputation.DprEarnings = newDprEarnings.String()

	// Save the updated reputation
	k.poaKeeper.SetReputation(ctx, reputation)

	log.Printf("Updated reputation saved: %+v\n", reputation)
	clientAccount, _ := sdk.AccAddressFromBech32(msg.Sender)
	errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, earningsCurrent)
	if errTransfer != nil {
		return nil, sdkerrors.Wrap(errTransfer, "[ClaimMotusRewards][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
	}

	// Save the updated DID back to the keeper
	k.didKeeper.SetClientDid(ctx, did)

	log.Println("############## Claim DPR Rewards Transaction Completed ##############")
	return &types.MsgClaimDprRewardsResponse{}, nil

}

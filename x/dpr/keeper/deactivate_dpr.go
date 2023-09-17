package keeper

import (
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Deactivate Dpr object
func (k Keeper) DeactivateDpr(ctx sdk.Context, totalEpoch uint64) {
	logger := k.Logger(ctx)
	activeDprs := k.GetAllActiveDpr(ctx)

	if activeDprs == nil {
		logger.Info("There is no active DPR to be deactivated.", "path", "DeactivateDpr")
		return
	}

	for _, dpr := range activeDprs {
		if totalEpoch > dpr.LengthOfDpr {
			newDpr := types.Dpr{
				Id:                            dpr.Id,
				Creator:                       dpr.Creator,
				PidSupportedOneToTwnety:       dpr.PidSupportedOneToTwnety,
				PidSupportedTwentyOneToForthy: dpr.PidSupportedTwentyOneToForthy,
				PidSupportedForthyOneToSixty:  dpr.PidSupportedForthyOneToSixty,
				IsActive:                      false,
				Vin:                           dpr.Vin,
				ClientPubkeys:                 dpr.ClientPubkeys,
				LengthOfDpr:                   dpr.LengthOfDpr,
			}
			k.SetDpr(ctx, newDpr)
		}
	}

	if logger != nil {
		logger.Info("Deactivation DPR successfully Done.", "path", "DeactivateDpr")
	}

}

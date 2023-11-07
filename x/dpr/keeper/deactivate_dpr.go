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
		if (dpr.Duration + dpr.DprStartEpoch) == totalEpoch {
			newDpr := types.Dpr{
				Id:             dpr.Id,
				Creator:        "",
				SupportedPIDs:  dpr.SupportedPIDs,
				IsActive:       false,
				Duration:       dpr.Duration,
				DprEndTime:     "",
				DprStartEpoch:  dpr.DprStartEpoch,
				DprBudget:      dpr.DprBudget,
				MaxClientCount: dpr.MaxClientCount,
			}
			k.SetDpr(ctx, newDpr)
		}
	}

	if logger != nil {
		logger.Info("Deactivation DPR successfully Done.", "path", "DeactivateDpr")
	}

}

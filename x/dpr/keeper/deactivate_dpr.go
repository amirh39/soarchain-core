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
				Creator:        dpr.Creator,
				SupportedPIDs:  dpr.SupportedPIDs,
				Status:         3,
				Duration:       dpr.Duration,
				DprEndTime:     dpr.DprEndTime,
				DprStartEpoch:  dpr.DprStartEpoch,
				DprBudget:      dpr.DprBudget,
				MaxClientCount: dpr.MaxClientCount,
				Name:           dpr.Name,
				ClientCounter:  dpr.ClientCounter,
			}
			k.SetDpr(ctx, newDpr)
		}
	}

	if logger != nil {
		logger.Info("Deactivation DPR successfully Done.", "path", "DeactivateDpr")
	}

}

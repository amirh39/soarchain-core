package keeper

import (
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Deactivate Dpr object
func (k Keeper) DeactivateDpr(ctx sdk.Context, totalEpoch uint64) {

	activeDprs := k.GetAllActiveDpr(ctx)

	for _, dpr := range activeDprs {
		if totalEpoch > dpr.LengthOfDpr {
			newDpr := types.Dpr{
				Id:                   dpr.Id,
				Creator:              dpr.Creator,
				PidSupported_1To_20:  dpr.PidSupported_1To_20,
				PidSupported_21To_40: dpr.PidSupported_21To_40,
				PidSupported_41To_60: dpr.PidSupported_41To_60,
				IsActive:             false,
				Vin:                  dpr.Vin,
				ClientPubkeys:        dpr.ClientPubkeys,
				LengthOfDpr:          dpr.LengthOfDpr,
			}
			k.SetDpr(ctx, newDpr)
		}
	}

}

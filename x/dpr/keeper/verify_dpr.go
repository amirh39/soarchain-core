package keeper

import "soarchain/x/dpr/types"

func (k Keeper) VerifyDprInputs(msg *types.MsgGenDpr, totalEpoch uint64) bool {
	if msg.Creator == "" {
		return false
	}

	if !msg.PidSupportedOneToTwnety && !msg.PidSupportedTwentyOneToForthy && !msg.PidSupportedForthyOneToSixty {
		return false
	}

	if msg.LengthOfDpr < totalEpoch {
		return false
	}

	return true
}

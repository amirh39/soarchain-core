package keeper

import (
	"soarchain/x/dpr/types"
	utility "soarchain/x/dpr/utility"
)

func (k Keeper) VerifyDprInputs(msg *types.MsgGenDpr, totalEpoch uint64) bool {
	if msg.Creator == "" {
		return false
	}
	if !(utility.IsValidHex(msg.SupportedPIDs)) {
		return false
	}

	return true
}

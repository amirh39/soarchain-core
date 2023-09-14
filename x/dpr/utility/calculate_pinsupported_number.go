package utility

import "soarchain/x/dpr/types"

func CalculatePinNumber(msg *types.MsgGenDpr) []uint {

	var pinRange []uint

	if msg.PidSupportedOneToTwnety {
		pinRange = append(pinRange, 1)
	}

	if msg.PidSupportedTwentyOneToForthy {
		pinRange = append(pinRange, 2)
	}

	if msg.GetPidSupportedForthyOneToSixty() {
		pinRange = append(pinRange, 3)
	}

	return pinRange

}

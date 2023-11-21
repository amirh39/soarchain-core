package keeper

import (
	"strings"

	"github.com/soar-robotics/soarchain-core/app/params"
	"github.com/soar-robotics/soarchain-core/x/dpr/types"
	"github.com/soar-robotics/soarchain-core/x/dpr/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) VerifyDprInputs(msg *types.MsgGenDpr) bool {
	return isValidCreator(msg) &&
		isValidSupportedPIDs(*msg.GetSupportedPIDs()) &&
		isValidDprBudget(msg) &&
		isValidMaxClientCount(msg) &&
		isValidName(msg)
}

func isValidDuration(msg *types.MsgGenDpr) bool {
	// Check for positive and non-zero(uint)
	return msg.Duration != 0
}

func isValidCreator(msg *types.MsgGenDpr) bool {
	// Check the creator(no need)
	return msg.Creator != ""
}

func isValidSupportedPIDs(supportedPIDs types.SupportedPIDs) bool {
	// Function to check if a string is a valid hex or empty.
	checkHexOrEmpty := func(s string) bool {
		return s == "" || utility.IsValidHex(s)
	}

	return checkHexOrEmpty(supportedPIDs.Pid_1To_20) &&
		checkHexOrEmpty(supportedPIDs.Pid_21To_40) &&
		checkHexOrEmpty(supportedPIDs.Pid_41To_60) &&
		checkHexOrEmpty(supportedPIDs.Pid_61To_80) &&
		checkHexOrEmpty(supportedPIDs.Pid_81To_A0) &&
		checkHexOrEmpty(supportedPIDs.Pid_A1To_C0) &&
		checkHexOrEmpty(supportedPIDs.Pid_C1To_E0) &&
		checkHexOrEmpty(supportedPIDs.Pid_SVCTo_9)
}

func isValidDprBudget(msg *types.MsgGenDpr) bool {
	// Parse the coin from the message
	coin, err := sdk.ParseCoinNormalized(msg.DprBudget)
	if err != nil {
		// If parsing failed, the coin is not valid
		return false
	}

	// Check if the denomination of the coin matches the BondDenom
	return coin.Denom == params.BondDenom
}

func isValidMaxClientCount(msg *types.MsgGenDpr) bool {
	// Negative or 0
	return msg.MaxClientCount != 0
}

func isValidName(msg *types.MsgGenDpr) bool {
	trimmedName := strings.TrimSpace(msg.Name)
	nameLength := len(trimmedName)
	return nameLength > 0 && nameLength <= 100
}

func (k Keeper) VerifyEnterDprInputs(msg *types.MsgEnterDpr) bool {

	return isValidSupportedPID(*msg.GetSupportedPIDs())
}

func isValidSupportedPID(supportedPIDs types.SupportedPIDs) bool {
	// Function to check if a string is a valid hex or empty.
	checkHexOrEmpty := func(s string) bool {
		return s == "" || utility.IsValidHex(s)
	}

	return checkHexOrEmpty(supportedPIDs.Pid_1To_20) &&
		checkHexOrEmpty(supportedPIDs.Pid_21To_40) &&
		checkHexOrEmpty(supportedPIDs.Pid_41To_60) &&
		checkHexOrEmpty(supportedPIDs.Pid_61To_80) &&
		checkHexOrEmpty(supportedPIDs.Pid_81To_A0) &&
		checkHexOrEmpty(supportedPIDs.Pid_A1To_C0) &&
		checkHexOrEmpty(supportedPIDs.Pid_C1To_E0) &&
		checkHexOrEmpty(supportedPIDs.Pid_SVCTo_9)
}

func IsCarSupportsDpr(carSupportedPIDs *types.SupportedPIDs, dprSupportedPIDs *types.SupportedPIDs) (bool, error) {
	// Check if the PID field in DPR is not an empty string, then check for support.
	checkSupport := func(carPID, dprPID string) (bool, error) {
		if dprPID == "" {
			// Consider an empty dprPID as always supported.
			return true, nil
		}
		return utility.ArePIDsSupported(carPID, dprPID)
	}

	pidFields := []struct {
		carPID string
		dprPID string
	}{
		{carSupportedPIDs.Pid_1To_20, dprSupportedPIDs.Pid_1To_20},
		{carSupportedPIDs.Pid_21To_40, dprSupportedPIDs.Pid_21To_40},
		{carSupportedPIDs.Pid_41To_60, dprSupportedPIDs.Pid_41To_60},
		{carSupportedPIDs.Pid_61To_80, dprSupportedPIDs.Pid_61To_80},
		{carSupportedPIDs.Pid_81To_A0, dprSupportedPIDs.Pid_81To_A0},
		{carSupportedPIDs.Pid_A1To_C0, dprSupportedPIDs.Pid_A1To_C0},
		{carSupportedPIDs.Pid_C1To_E0, dprSupportedPIDs.Pid_C1To_E0},
		{carSupportedPIDs.Pid_SVCTo_9, dprSupportedPIDs.Pid_SVCTo_9},
	}

	for _, field := range pidFields {
		supported, err := checkSupport(field.carPID, field.dprPID)
		if err != nil {
			return false, err
		}
		if !supported {
			// The car does not support this particular PID, return false.
			return false, nil
		}
	}

	// All PIDs are supported by the car, return true.
	return true, nil
}

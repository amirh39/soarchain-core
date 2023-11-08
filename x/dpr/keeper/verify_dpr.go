package keeper

import (
	"soarchain/app/params"
	"soarchain/x/dpr/types"
	"soarchain/x/dpr/utility"
	"strings"

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
	return utility.IsValidHex(supportedPIDs.Pid_1To_20) &&
		utility.IsValidHex(supportedPIDs.Pid_21To_40) &&
		utility.IsValidHex(supportedPIDs.Pid_41To_60) &&
		utility.IsValidHex(supportedPIDs.Pid_61To_80) &&
		utility.IsValidHex(supportedPIDs.Pid_81To_A0) &&
		utility.IsValidHex(supportedPIDs.Pid_A1To_C0) &&
		utility.IsValidHex(supportedPIDs.Pid_C1To_E0) &&
		utility.IsValidHex(supportedPIDs.Pid_SVCTo_9)
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
	return utility.IsValidHex(supportedPIDs.Pid_1To_20) &&
		utility.IsValidHex(supportedPIDs.Pid_21To_40) &&
		utility.IsValidHex(supportedPIDs.Pid_41To_60) &&
		utility.IsValidHex(supportedPIDs.Pid_61To_80) &&
		utility.IsValidHex(supportedPIDs.Pid_81To_A0) &&
		utility.IsValidHex(supportedPIDs.Pid_A1To_C0) &&
		utility.IsValidHex(supportedPIDs.Pid_C1To_E0) &&
		utility.IsValidHex(supportedPIDs.Pid_SVCTo_9)
}

func IsCarSupportsDpr(carSupportedPIDs *types.SupportedPIDs, dprSupportedPIDs *types.SupportedPIDs) (bool, error) {
	carSupportsPID1To20, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_1To_20, dprSupportedPIDs.Pid_1To_20)
	if err != nil {
		return false, err
	}

	carSupportsPID21To40, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_21To_40, dprSupportedPIDs.Pid_21To_40)
	if err != nil {
		return false, err
	}

	carSupportsPID41To60, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_41To_60, dprSupportedPIDs.Pid_41To_60)
	if err != nil {
		return false, err
	}

	carSupportsPID61To80, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_61To_80, dprSupportedPIDs.Pid_61To_80)
	if err != nil {
		return false, err
	}

	carSupportsPID81ToA0, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_81To_A0, dprSupportedPIDs.Pid_81To_A0)
	if err != nil {
		return false, err
	}

	carSupportsPIDA1ToC0, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_A1To_C0, dprSupportedPIDs.Pid_A1To_C0)
	if err != nil {
		return false, err
	}

	carSupportsPIDC1ToE0, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_C1To_E0, dprSupportedPIDs.Pid_C1To_E0)
	if err != nil {
		return false, err
	}

	carSupportsPIDSvcTo9, err := utility.ArePIDsSupported(carSupportedPIDs.Pid_SVCTo_9, dprSupportedPIDs.Pid_SVCTo_9)
	if err != nil {
		return false, err
	}

	// The car supports the DPR if all required PIDs are supported.
	isSupported := carSupportsPID1To20 && carSupportsPID21To40 && carSupportsPID41To60 && carSupportsPID61To80 &&
		carSupportsPID81ToA0 && carSupportsPIDA1ToC0 && carSupportsPIDC1ToE0 && carSupportsPIDSvcTo9

	return isSupported, nil
}

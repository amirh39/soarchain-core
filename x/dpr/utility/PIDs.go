package utility

import (
	"fmt"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func hexToBin(hexValue string) (string, error) {
	i, err := strconv.ParseInt(hexValue, 16, 64)
	if err != nil {
		return "", sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[hexToBin] Error parsing hex value: "+err.Error())
	}
	return fmt.Sprintf("%032s", strconv.FormatInt(i, 2)), nil
}

func getAvailablePIDs(binaryValue string) []string {
	var availablePIDs []string

	for i, bit := range binaryValue {
		if bit == '1' {
			// PIDs are 1-based, so add 1 to the index and format as a two-digit string
			pid := fmt.Sprintf("%02d", i+1)
			availablePIDs = append(availablePIDs, pid)
		}
	}

	return availablePIDs
}

func ArePIDsSupported(carHex, dprHex string) (bool, error) {
	carBinary, err := hexToBin(carHex)
	if err != nil {
		return false, err
	}

	dprBinary, err := hexToBin(dprHex)
	if err != nil {
		return false, err
	}

	carPIDs := getAvailablePIDs(carBinary)
	dprPIDs := getAvailablePIDs(dprBinary)

	// Check if all PIDs in the DPR are supported by the car
	for _, dprPID := range dprPIDs {
		if !contains(carPIDs, dprPID) {
			return false, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[arePIDsSupported] Car does not support PID: "+dprPID)
		}
	}

	return true, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

package utility

import "soarchain/x/poa/utility/utilConstants"

var addressLength = utilConstants.AddressLength

/** ValidString performs the validation of a string input */
func ValidString(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}

package utility

import "soarchain/x/poa/utility/utilConstants"

var addressLength = utilConstants.AddressLength

/** ValidPublickey performs the validation of a publickey */
func ValidPubkey(input string) bool { // TODO: Add validation of an object Pubkey
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}

/** ValidAddress performs the validation of an address */
func ValidAddress(input string) bool {
	if len(input) > addressLength || len(input) < addressLength || input == "" || input[0:4] != "soar" {
		return false
	}
	return true
}

/** ValidString performs the validation of a string input */
func ValidString(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}

package utility

/** IsValidPublickey performs the validation of publickey */
func ValidPubkey(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}

/** IsValidAddress performs the validation of address */
func IsValidAddress(input string) bool {
	if len(input) > 43 || len(input) < 43 || input == "" || input[0:4] != "soar" {
		return false
	}
	return true
}

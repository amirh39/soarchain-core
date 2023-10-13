package utility

import "regexp"

func IsValidHex(hexValue string) bool {
	// Use a regular expression to match valid hexadecimal characters (0-9, A-F).
	validHexPattern := "^[0-9A-Fa-f]+$"
	validHex := regexp.MustCompile(validHexPattern)

	return validHex.MatchString(hexValue)
}

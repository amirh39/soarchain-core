package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidateDidAddress(address string) bool {
	_, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return false
	}
	return true
}

func ValidateStakeAmount(amount string) bool {
	_, err := sdk.ParseCoinNormalized(amount)
	if err != nil {
		return false
	}
	return true
}

func ValidateSupportedPIDs(supportedPIDs string) bool {
	// Use a regular expression to match valid hexadecimal characters (0-9, A-F).
	pattern := "^[0-9A-Fa-f]+$"
	matched := regexp.MustCompile(pattern)
	if !matched.MatchString(supportedPIDs) {
		return false
	}
	return true
}

package utility

import (
	"soarchain/x/poa/utility/utilConstants"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var addressLength = utilConstants.AddressLength

/** ValidString performs the validation of a string input */
func ValidString(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}

func ValidateRewardAmount(amount string) bool {
	_, err := sdk.ParseCoinNormalized(amount)
	if err != nil {
		return false
	}
	return true
}

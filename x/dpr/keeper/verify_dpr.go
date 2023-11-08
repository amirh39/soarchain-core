package keeper

import (
	"soarchain/app/params"
	"soarchain/x/dpr/types"
	utility "soarchain/x/dpr/utility"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) VerifyDprInputs(msg *types.MsgGenDpr) bool {
	return isValidCreator(msg) &&
		isValidSupportedPIDs(msg) &&
		isValidDprBudget(msg) &&
		isValidMaxClientCount(msg) &&
		isValidName(msg)
}

func isValidCreator(msg *types.MsgGenDpr) bool {
	// Check the creator(no need)
	return msg.Creator != ""
}

func isValidSupportedPIDs(msg *types.MsgGenDpr) bool {
	// Check if hex is valid using regex
	return utility.IsValidHex(msg.SupportedPIDs)
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

	return isValidSupportedPID(msg)
}

func isValidSupportedPID(msg *types.MsgEnterDpr) bool {
	// Check if hex is valid using regex
	return utility.IsValidHex(msg.SupportedPIDs)
}

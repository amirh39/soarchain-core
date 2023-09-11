package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// FindAccount find a specific address from an account list
func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, bool) {
	creator, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[AccAddressFromBech32] failed. Empty address string is not allowed.")
		return simtypes.Account{}, false
	}
	return simtypes.FindAccount(accs, creator)
}

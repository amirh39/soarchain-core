package errors

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ConverToString    = "[Type Error] Couldn't convert from Int to String."
	ParseUint         = "Couldn't parse data."
	ErrInvalidType    = "Couln't find valid type"
	ErrInvalidAddress = "[AccAddressFromBech32] failed. Empty address string is not allowed."
	ErrDidNotFound    = "[Did][GetDidDocument] failed. Did not found."
	InvalidRequest    = "Invalid Request."
	ErrDidNotActive   = "Did is Deactivated."
)

var (
	ErrInvalidDid = sdkerrors.Register("ModuleName", 3, "Invalid Did")
)

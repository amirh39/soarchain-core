package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnregisterGuard = "unregister_guard"

var _ sdk.Msg = &MsgUnregisterGuard{}

func NewMsgUnregisterGuard(creator string) *MsgUnregisterGuard {
	return &MsgUnregisterGuard{
		Creator: creator,
	}
}

func (msg *MsgUnregisterGuard) Route() string {
	return RouterKey
}

func (msg *MsgUnregisterGuard) Type() string {
	return TypeMsgUnregisterGuard
}

func (msg *MsgUnregisterGuard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnregisterGuard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnregisterGuard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

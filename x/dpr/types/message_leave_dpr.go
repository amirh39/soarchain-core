package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLeaveDpr = "leave_dpr"

var _ sdk.Msg = &MsgActivateDpr{}

func NewMsgLeaveDpr(sender string, dprId string) *MsgLeaveDpr {
	return &MsgLeaveDpr{
		Sender: sender,
		DprId:  dprId,
	}
}

func (msg *MsgLeaveDpr) Route() string {
	return RouterKey
}

func (msg *MsgLeaveDpr) Type() string {
	return TypeMsgLeaveDpr
}

func (msg *MsgLeaveDpr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgLeaveDpr][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLeaveDpr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLeaveDpr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgLeaveDpr][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

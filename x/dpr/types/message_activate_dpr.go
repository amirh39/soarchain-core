package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgActivateDpr = "activate_dpr"

var _ sdk.Msg = &MsgActivateDpr{}

func NewMsgActivateDpr(sender string, dprId string) *MsgActivateDpr {
	return &MsgActivateDpr{
		Sender: sender,
		DprId:  dprId,
	}
}

func (msg *MsgActivateDpr) Route() string {
	return RouterKey
}

func (msg *MsgActivateDpr) Type() string {
	return TypeMsgActivateDpr
}

func (msg *MsgActivateDpr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgActivateDpr][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgActivateDpr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgActivateDpr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgActivateDpr][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

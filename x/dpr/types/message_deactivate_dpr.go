package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeactivateDpr = "deactivate_dpr"

var _ sdk.Msg = &MsgActivateDpr{}

func NewMsgDeactivateDpr(sender string, dprId string, duration uint64) *MsgDeactivateDpr {
	return &MsgDeactivateDpr{
		Sender:   sender,
		DprId:    dprId,
		Duration: duration,
	}
}

func (msg *MsgDeactivateDpr) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateDpr) Type() string {
	return TypeMsgActivateDpr
}

func (msg *MsgDeactivateDpr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgDeactivateDpr][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeactivateDpr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateDpr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgDeactivateDpr][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

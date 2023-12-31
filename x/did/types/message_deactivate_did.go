package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeactivateDid = "deactivate_did"

var _ sdk.Msg = &MsgDeactivateDid{}

func NewMsgDeactivateDid(creator string) *MsgDeactivateDid {
	return &MsgDeactivateDid{
		Creator: creator,
	}
}

func (msg *MsgDeactivateDid) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateDid) Type() string {
	return TypeMsgDeactivateDid
}

func (msg *MsgDeactivateDid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgDeactivateDid][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeactivateDid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateDid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgDeactivateDid][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

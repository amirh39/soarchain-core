package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnterDpr = "gen_enter"

var _ sdk.Msg = &MsgGenDpr{}

func NewMsgEnterDpr(sender string, pubKey string, dprId string) *MsgEnterDpr {
	return &MsgEnterDpr{
		PubKey: pubKey,
		Sender: sender,
		DprId:  dprId,
	}
}

func (msg *MsgEnterDpr) Route() string {
	return RouterKey
}

func (msg *MsgEnterDpr) Type() string {
	return TypeMsgEnterDpr
}

func (msg *MsgEnterDpr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgEnterDpr][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnterDpr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnterDpr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgEnterDpr][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

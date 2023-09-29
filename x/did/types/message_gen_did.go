package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenDid = "gen_did"

var _ sdk.Msg = &MsgGenDid{}

func NewMsgGenDid(document DidDocument, signature string, certificate string, creator string) *MsgGenDid {
	return &MsgGenDid{
		Document:    &document,
		Signature:   signature,
		Certificate: certificate,
		Creator:     creator,
	}
}

func (msg *MsgGenDid) Route() string {
	return RouterKey
}

func (msg *MsgGenDid) Type() string {
	return TypeMsgGenDid
}

func (msg *MsgGenDid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgGenDid][AccAddressFromBech32] failed. Empty creator address is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenDid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenDid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgGenDid][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

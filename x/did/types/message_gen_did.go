package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenDid = "gen_did"

var _ sdk.Msg = &MsgGenDid{}

func NewMsgGenDid(did string, document DidDocument, verificationMethodID string, signature []byte, fromAddress string) *MsgGenDid {
	return &MsgGenDid{
		Document:  &document,
		Signature: signature,
		Creator:   fromAddress,
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
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgGenDid][AccAddressFromBech32] failed. Empty address string is not allowed.")
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

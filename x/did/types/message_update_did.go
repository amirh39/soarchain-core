package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDid = "update_did"

var _ sdk.Msg = &MsgUpdateDid{}

func NewMsgUpdateDid(did string, document ClientDidDocument, verificationMethodID string, signature []byte, fromAddress string) *MsgUpdateDid {
	return &MsgUpdateDid{
		Did:                  did,
		Document:             &document,
		VerificationMethodId: verificationMethodID,
		Signature:            signature,
		FromAddress:          fromAddress,
	}
}

func (msg *MsgUpdateDid) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDid) Type() string {
	return TypeMsgUpdateDid
}

func (msg *MsgUpdateDid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgUpdateDid][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgUpdateDid][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

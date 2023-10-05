package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenClient = "gen_client"

var _ sdk.Msg = &MsgGenClient{}

func NewMsgGenClient(document ClientDidDocument, signature string, certificate string, creator string) *MsgGenClient {
	return &MsgGenClient{
		Document:    &document,
		Signature:   signature,
		Certificate: certificate,
		Creator:     creator,
	}
}

func (msg *MsgGenClient) Route() string {
	return RouterKey
}

func (msg *MsgGenClient) Type() string {
	return TypeMsgGenClient
}

func (msg *MsgGenClient) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgGenClient][AccAddressFromBech32] failed. Empty creator address is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenClient) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgGenClient][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenClient = "gen_client"

var _ sdk.Msg = &MsgGenClient{}

func NewMsgGenClient(creator string, certificate string) *MsgGenClient {
	return &MsgGenClient{
		Creator:     creator,
		Certificate: certificate,
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
		panic(err)
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

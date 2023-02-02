package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnregisterClient = "unregister_client"

var _ sdk.Msg = &MsgUnregisterClient{}

func NewMsgUnregisterClient(creator string, pubkey string, fee string) *MsgUnregisterClient {
	return &MsgUnregisterClient{
		Creator: creator,
		Pubkey:  pubkey,
		Fee:     fee,
	}
}

func (msg *MsgUnregisterClient) Route() string {
	return RouterKey
}

func (msg *MsgUnregisterClient) Type() string {
	return TypeMsgUnregisterClient
}

func (msg *MsgUnregisterClient) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnregisterClient) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnregisterClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

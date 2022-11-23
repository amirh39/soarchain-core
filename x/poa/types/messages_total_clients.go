package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTotalClients = "create_total_clients"
	TypeMsgUpdateTotalClients = "update_total_clients"
	TypeMsgDeleteTotalClients = "delete_total_clients"
)

var _ sdk.Msg = &MsgCreateTotalClients{}

func NewMsgCreateTotalClients(creator string, count uint64) *MsgCreateTotalClients {
	return &MsgCreateTotalClients{
		Creator: creator,
		Count:   count,
	}
}

func (msg *MsgCreateTotalClients) Route() string {
	return RouterKey
}

func (msg *MsgCreateTotalClients) Type() string {
	return TypeMsgCreateTotalClients
}

func (msg *MsgCreateTotalClients) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTotalClients) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTotalClients) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTotalClients{}

func NewMsgUpdateTotalClients(creator string, count uint64) *MsgUpdateTotalClients {
	return &MsgUpdateTotalClients{
		Creator: creator,
		Count:   count,
	}
}

func (msg *MsgUpdateTotalClients) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTotalClients) Type() string {
	return TypeMsgUpdateTotalClients
}

func (msg *MsgUpdateTotalClients) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTotalClients) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTotalClients) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTotalClients{}

func NewMsgDeleteTotalClients(creator string) *MsgDeleteTotalClients {
	return &MsgDeleteTotalClients{
		Creator: creator,
	}
}
func (msg *MsgDeleteTotalClients) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTotalClients) Type() string {
	return TypeMsgDeleteTotalClients
}

func (msg *MsgDeleteTotalClients) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTotalClients) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTotalClients) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

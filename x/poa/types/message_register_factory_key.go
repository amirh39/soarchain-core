package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterFactoryKey = "register_factory_key"

var _ sdk.Msg = &MsgRegisterFactoryKey{}

func NewMsgRegisterFactoryKey(creator string, factoryKey string) *MsgRegisterFactoryKey {
	return &MsgRegisterFactoryKey{
		Creator:    creator,
		FactoryKey: factoryKey,
	}
}

func (msg *MsgRegisterFactoryKey) Route() string {
	return RouterKey
}

func (msg *MsgRegisterFactoryKey) Type() string {
	return TypeMsgRegisterFactoryKey
}

func (msg *MsgRegisterFactoryKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterFactoryKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterFactoryKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

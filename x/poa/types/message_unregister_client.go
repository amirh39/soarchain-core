package types

import (
	"soarchain/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnregisterClient = "unregister_client"

var _ sdk.Msg = &MsgUnregisterClient{}

func NewMsgUnregisterClient(creator string, pubkey string) *MsgUnregisterClient {
	return &MsgUnregisterClient{
		Creator: creator,
		Pubkey:  pubkey,
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
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
		return nil
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

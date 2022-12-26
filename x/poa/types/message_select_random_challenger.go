package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSelectRandomChallenger = "select_random_challenger"

var _ sdk.Msg = &MsgSelectRandomChallenger{}

func NewMsgSelectRandomChallenger(creator string, multiplier string) *MsgSelectRandomChallenger {
	return &MsgSelectRandomChallenger{
		Creator:    creator,
		Multiplier: multiplier,
	}
}

func (msg *MsgSelectRandomChallenger) Route() string {
	return RouterKey
}

func (msg *MsgSelectRandomChallenger) Type() string {
	return TypeMsgSelectRandomChallenger
}

func (msg *MsgSelectRandomChallenger) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSelectRandomChallenger) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSelectRandomChallenger) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

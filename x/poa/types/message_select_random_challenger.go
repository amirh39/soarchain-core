package types

import (
	"soarchain/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSelectRandomChallenger = "select_random_challenger"

var _ sdk.Msg = &MsgSelectRandomChallenger{}

func NewMsgSelectRandomChallenger(creator string) *MsgSelectRandomChallenger {
	return &MsgSelectRandomChallenger{
		Creator: creator,
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
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
		return nil
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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenChallenger = "gen_challenger"

var _ sdk.Msg = &MsgGenChallenger{}

func NewMsgGenChallenger(document ChallengerDid, signature string, certificate string, creator string) *MsgGenChallenger {
	return &MsgGenChallenger{
		Document:    &document,
		Signature:   signature,
		Certificate: certificate,
		Creator:     creator,
	}
}

func (msg *MsgGenChallenger) Route() string {
	return RouterKey
}

func (msg *MsgGenChallenger) Type() string {
	return TypeMsgGenClient
}

func (msg *MsgGenChallenger) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgGenChallenger][AccAddressFromBech32] failed. Empty creator address is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenChallenger) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenChallenger) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgGenChallenger][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenChallenger = "gen_challenger"

var _ sdk.Msg = &MsgGenChallenger{}

func NewMsgGenChallenger(creator string, challengerStake string, challengerIp string, challengerType string, certificate string, signature string) *MsgGenChallenger {
	return &MsgGenChallenger{
		Creator:         creator,
		ChallengerStake: challengerStake,
		ChallengerIp:    challengerIp,
		ChallengerType:  challengerType,
		Certificate:     certificate,
		Signature:       signature,
	}
}

func (msg *MsgGenChallenger) Route() string {
	return RouterKey
}

func (msg *MsgGenChallenger) Type() string {
	return TypeMsgGenChallenger
}

func (msg *MsgGenChallenger) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgGenChallenger][AccAddressFromBech32] failed. Empty address string is not allowed.")
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

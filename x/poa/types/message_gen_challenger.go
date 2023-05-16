package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenChallenger = "gen_challenger"

var _ sdk.Msg = &MsgGenChallenger{}

func NewMsgGenChallenger(creator string, challengerPubKey string, challengerAddr string, challengerStake string, challengerIp string, challengerType string, certificate string, signature string) *MsgGenChallenger {
	return &MsgGenChallenger{
		Creator:          creator,
		ChallengerPubKey: challengerPubKey,
		ChallengerAddr:   challengerAddr,
		ChallengerStake:  challengerStake,
		ChallengerIp:     challengerIp,
		Challengertype:   challengerType,
		Certificate:      certificate,
		Signature:        signature,
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
		panic(err)
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

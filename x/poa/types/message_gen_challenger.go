package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenChallenger = "gen_challenger"

var _ sdk.Msg = &MsgGenChallenger{}

func NewMsgGenChallenger(creator string, stakeAmount string) *MsgGenChallenger {
	return &MsgGenChallenger{
		Creator:     creator,
		StakeAmount: stakeAmount,
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

package types

import (
	"soarchain/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnregisterChallenger = "unregister_challenger"

var _ sdk.Msg = &MsgUnregisterChallenger{}

func NewMsgUnregisterChallenger(creator string, challengerAddress string) *MsgUnregisterChallenger {
	return &MsgUnregisterChallenger{
		Creator:           creator,
		ChallengerAddress: challengerAddress,
	}
}

func (msg *MsgUnregisterChallenger) Route() string {
	return RouterKey
}

func (msg *MsgUnregisterChallenger) Type() string {
	return TypeMsgUnregisterChallenger
}

func (msg *MsgUnregisterChallenger) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnregisterChallenger) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnregisterChallenger) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

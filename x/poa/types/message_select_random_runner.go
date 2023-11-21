package types

import (
	"github.com/soar-robotics/soarchain-core/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSelectRandomRunner = "select_random_runner"

var _ sdk.Msg = &MsgSelectRandomRunner{}

func NewMsgSelectRandomRunner(creator string) *MsgSelectRandomRunner {
	return &MsgSelectRandomRunner{
		Creator: creator,
	}
}

func (msg *MsgSelectRandomRunner) Route() string {
	return RouterKey
}

func (msg *MsgSelectRandomRunner) Type() string {
	return TypeMsgSelectRandomRunner
}

func (msg *MsgSelectRandomRunner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSelectRandomRunner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSelectRandomRunner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgSelectRandomRunner] failed. Invalid creator address (%s)", msg.Creator)
	}
	return nil
}

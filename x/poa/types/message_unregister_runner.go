package types

import (
	"soarchain/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnregisterRunner = "unregister_runner"

var _ sdk.Msg = &MsgUnregisterRunner{}

func NewMsgUnregisterRunner(creator string, runnerAddress string) *MsgUnregisterRunner {
	return &MsgUnregisterRunner{
		Creator:       creator,
		RunnerAddress: runnerAddress,
	}
}

func (msg *MsgUnregisterRunner) Route() string {
	return RouterKey
}

func (msg *MsgUnregisterRunner) Type() string {
	return TypeMsgUnregisterRunner
}

func (msg *MsgUnregisterRunner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnregisterRunner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnregisterRunner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

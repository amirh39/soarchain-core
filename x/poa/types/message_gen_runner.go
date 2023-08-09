package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenRunner = "gen_runner"

var _ sdk.Msg = &MsgGenRunner{}

func NewMsgGenRunner(creator string, runnerStake string, runnerIpAddress string, certificate string, signature string) *MsgGenRunner {
	return &MsgGenRunner{
		Creator:     creator,
		RunnerStake: runnerStake,
		RunnerIp:    runnerIpAddress,
		Certificate: certificate,
		Signature:   signature,
	}
}

func (msg *MsgGenRunner) Route() string {
	return RouterKey
}

func (msg *MsgGenRunner) Type() string {
	return TypeMsgGenRunner
}

func (msg *MsgGenRunner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgGenRunner][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenRunner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenRunner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgGenRunner][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

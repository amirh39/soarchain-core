package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenrunner = "gen_runner"

var _ sdk.Msg = &MsgGenRunner{}

func NewMsgGenRunner(creator string, runnerPubKey string, runnerAddress string, runnerStake string, runnerIpaddress string) *MsgGenRunner {
	return &MsgGenRunner{
		Creator:      creator,
		RunnerPubKey: runnerPubKey,
		RunnerAddr:   runnerAddress,
		RunnerStake:  runnerStake,
		RunnerIp:     runnerIpaddress,
	}
}

func (msg *MsgGenRunner) Route() string {
	return RouterKey
}

func (msg *MsgGenRunner) Type() string {
	return TypeMsgGenrunner
}

func (msg *MsgGenRunner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

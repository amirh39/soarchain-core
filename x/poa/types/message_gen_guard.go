package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenGuard = "gen_guard"

var _ sdk.Msg = &MsgGenGuard{}

func NewMsgGenGuard(creator string, guardPubKey string, v2XAddr string, v2XStake string, v2XIp string, v2NAddr string, v2NStake string, v2NIp string, runnerAddr string, runnerStake string, runnerIp string) *MsgGenGuard {
	return &MsgGenGuard{
		Creator:     creator,
		GuardPubKey: guardPubKey,
		V2XAddr:     v2XAddr,
		V2XStake:    v2XStake,
		V2XIp:       v2XIp,
		V2NAddr:     v2NAddr,
		V2NStake:    v2NStake,
		V2NIp:       v2NIp,
		RunnerAddr:  runnerAddr,
		RunnerStake: runnerStake,
		RunnerIp:    runnerIp,
	}
}

func (msg *MsgGenGuard) Route() string {
	return RouterKey
}

func (msg *MsgGenGuard) Type() string {
	return TypeMsgGenGuard
}

func (msg *MsgGenGuard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenGuard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenGuard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

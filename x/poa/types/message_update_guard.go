package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateGuard = "update_guard"

var _ sdk.Msg = &MsgUpdateGuard{}

func NewMsgUpdateGuard(creator string, v2XPubKey string, v2XAddr string, v2XStake string, v2XIp string, v2NPubKey string, v2NAddr string, v2NStake string, v2NIp string, runnerPubKey string, runnerAddr string, runnerStake string, runnerIp string) *MsgGenGuard {
	return &MsgGenGuard{
		Creator:      creator,
		V2XPubKey:    v2XPubKey,
		V2XAddr:      v2XAddr,
		V2XStake:     v2XStake,
		V2XIp:        v2XIp,
		V2NPubKey:    v2NPubKey,
		V2NAddr:      v2NAddr,
		V2NStake:     v2NStake,
		V2NIp:        v2NIp,
		RunnerPubKey: runnerPubKey,
		RunnerAddr:   runnerAddr,
		RunnerStake:  runnerStake,
		RunnerIp:     runnerIp,
	}
}

func (msg *MsgUpdateGuard) Route() string {
	return RouterKey
}

func (msg *MsgUpdateGuard) Type() string {
	return TypeMsgUpdateGuard
}

func (msg *MsgUpdateGuard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateGuard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateGuard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

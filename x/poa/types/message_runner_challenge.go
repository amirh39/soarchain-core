package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRunnerChallenge = "runner_challenge"

var _ sdk.Msg = &MsgRunnerChallenge{}

func NewMsgRunnerChallenge(creator string, runnerAddress string, v2nDeviceType string, challengeResult string) *MsgRunnerChallenge {
	return &MsgRunnerChallenge{
		Creator:         creator,
		RunnerAddress:   runnerAddress,
		V2NDeviceType:   v2nDeviceType,
		ChallengeResult: challengeResult,
	}
}

func (msg *MsgRunnerChallenge) Route() string {
	return RouterKey
}

func (msg *MsgRunnerChallenge) Type() string {
	return TypeMsgRunnerChallenge
}

func (msg *MsgRunnerChallenge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRunnerChallenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRunnerChallenge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

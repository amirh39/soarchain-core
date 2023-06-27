package types

import (
	"soarchain/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRunnerChallenge = "runner_challenge"

var _ sdk.Msg = &MsgRunnerChallenge{}

func NewMsgRunnerChallenge(creator string, runnerPubkey string, clientPubkeys []string, challengeResult string) *MsgRunnerChallenge {
	return &MsgRunnerChallenge{
		Creator:         creator,
		RunnerpubKey:    runnerPubkey,
		ClientPubkeys:   clientPubkeys,
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
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
		return nil
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

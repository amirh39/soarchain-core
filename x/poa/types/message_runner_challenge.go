package types

import (
	"strings"

	"github.com/soar-robotics/soarchain-core/x/poa/constants"
	"github.com/soar-robotics/soarchain-core/x/poa/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRunnerChallenge = "runner_challenge"

var _ sdk.Msg = &MsgRunnerChallenge{}

func NewMsgRunnerChallenge(creator string, runnerPubkey string, clientPubkeys []string, challengeResult string) *MsgRunnerChallenge {
	return &MsgRunnerChallenge{
		Creator:         creator,
		RunnerPubkey:    runnerPubkey,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgRunnerChallenge][ValidateBasic] failed. Invalid creator address [ %s ] ,", msg.Creator)
	}
	if msg.RunnerPubkey == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[NewMsgRunnerChallenge][ValidateBasic] failed. Invalid runner pubkey [ %s ] ", msg.RunnerPubkey)
	}
	if len(msg.ClientPubkeys) < 1 || msg.ClientPubkeys == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[NewMsgRunnerChallenge][ValidateBasic] failed. Invalid client pubkey [ %s ] ", msg.ClientPubkeys)
	}

	reward := strings.Compare(msg.ChallengeResult, constants.Reward)
	punish := strings.Compare(msg.ChallengeResult, constants.Punish)

	if reward != 0 && punish != 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[NewMsgRunnerChallenge][ValidateBasic] failed. Invalid challenge result [ %s ] ", msg.ChallengeResult)
	}

	return nil
}

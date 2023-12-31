package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChallengeService = "challenge_service"

var _ sdk.Msg = &MsgChallengeService{}

func NewMsgChallengeService(creator string, clientPubkey string, clientCommunicationMode string, challengeResult string) *MsgChallengeService {
	return &MsgChallengeService{
		Creator:                 creator,
		ClientPubkey:            clientPubkey,
		ClientCommunicationMode: clientCommunicationMode,
		ChallengeResult:         challengeResult,
	}
}

func (msg *MsgChallengeService) Route() string {
	return RouterKey
}

func (msg *MsgChallengeService) Type() string {
	return TypeMsgChallengeService
}

func (msg *MsgChallengeService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgChallengeService][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChallengeService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChallengeService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgChallengeService][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

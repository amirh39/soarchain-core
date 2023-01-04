package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgV2NChallenge = "v_2_n_challenge"

var _ sdk.Msg = &MsgV2NChallenge{}

func NewMsgV2NChallenge(creator string, runnerAddress string, runnerResult string, v2NBxAddress []string, v2NBxResult []string) *MsgV2NChallenge {
	return &MsgV2NChallenge{
		Creator:       creator,
		RunnerAddress: runnerAddress,
		RunnerResult:  runnerResult,
		V2NBxAddress:  v2NBxAddress,
		V2NBxResult:   v2NBxResult,
	}
}

func (msg *MsgV2NChallenge) Route() string {
	return RouterKey
}

func (msg *MsgV2NChallenge) Type() string {
	return TypeMsgV2NChallenge
}

func (msg *MsgV2NChallenge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgV2NChallenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgV2NChallenge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

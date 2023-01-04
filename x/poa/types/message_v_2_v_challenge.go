package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgV2VChallenge = "v_2_v_challenge"

var _ sdk.Msg = &MsgV2VChallenge{}

func NewMsgV2VChallenge(creator string, rxAddress string, rxResult string, bxAddress []string, bxResult []string) *MsgV2VChallenge {
	return &MsgV2VChallenge{
		Creator:   creator,
		RxAddress: rxAddress,
		RxResult:  rxResult,
		BxAddress: bxAddress,
		BxResult:  bxResult,
	}
}

func (msg *MsgV2VChallenge) Route() string {
	return RouterKey
}

func (msg *MsgV2VChallenge) Type() string {
	return TypeMsgV2VChallenge
}

func (msg *MsgV2VChallenge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgV2VChallenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgV2VChallenge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDpr = "update_dpr"

var _ sdk.Msg = &MsgUpdateDpr{}

func NewMsgUpdateDpr(DprId string, duration uint64, MaxClientCount uint64, DprBudget string, sender string) *MsgUpdateDpr {
	return &MsgUpdateDpr{
		DprId:          DprId,
		Duration:       duration,
		MaxClientCount: MaxClientCount,
		DprBudget:      DprBudget,
		Sender:         sender,
	}
}

func (msg *MsgUpdateDpr) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDpr) Type() string {
	return TypeMsgGenDpr
}

func (msg *MsgUpdateDpr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[NewMsgUpdateDpr][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDpr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDpr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgUpdateDpr][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

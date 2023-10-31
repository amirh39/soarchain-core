package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenDpr = "gen_dpr"

var _ sdk.Msg = &MsgGenDpr{}

func NewMsgGenDpr(creator string, SupportedPIDs string, duration uint64, dprBudget string, maxClientCount uint64) *MsgGenDpr {
	return &MsgGenDpr{
		Creator:        creator,
		SupportedPIDs:  SupportedPIDs,
		Duration:       duration,
		DprBudget:      dprBudget,
		MaxClientCount: maxClientCount,
	}
}

func (msg *MsgGenDpr) Route() string {
	return RouterKey
}

func (msg *MsgGenDpr) Type() string {
	return TypeMsgGenDpr
}

func (msg *MsgGenDpr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgGenDpr][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenDpr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenDpr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgGenDpr][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

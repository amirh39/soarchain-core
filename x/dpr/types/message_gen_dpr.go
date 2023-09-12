package types

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenDpr = "gen_dpr"

var _ sdk.Msg = &MsgGenDpr{}

func NewMsgGenDpr(creator string, pidSupported_1_to_20 bool, pidSupported_21_to_40 bool, pidSupported_41_to_60 bool, vin []string, lengthOfDpr uint64) *MsgGenDpr {
	log.Println("444444444444444444444")
	return &MsgGenDpr{
		Creator:              creator,
		PidSupported_1To_20:  pidSupported_1_to_20,
		PidSupported_21To_40: pidSupported_21_to_40,
		PidSupported_41To_60: pidSupported_41_to_60,
		Vin:                  vin,
		LengthOfDpr:          lengthOfDpr,
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

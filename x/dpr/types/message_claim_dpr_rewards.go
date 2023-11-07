package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimDprRewards = "claim_dpr_rewards"

var _ sdk.Msg = &MsgClaimDprRewards{}

func NewMsgClaimDprRewards(sender string) *MsgClaimDprRewards {
	return &MsgClaimDprRewards{
		Sender: sender,
	}
}

func (msg *MsgClaimDprRewards) Route() string {
	return RouterKey
}

func (msg *MsgClaimDprRewards) Type() string {
	return TypeMsgActivateDpr
}

func (msg *MsgClaimDprRewards) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[MsgClaimDprRewards][AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimDprRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimDprRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[MsgClaimDprRewards][ValidateBasic] failed. Invalid creator address (%s)", err)
	}
	return nil
}

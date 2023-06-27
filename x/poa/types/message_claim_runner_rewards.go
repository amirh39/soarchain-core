package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimRunnerRewards = "claim_runner_rewards"

var _ sdk.Msg = &MsgClaimRunnerRewards{}

func NewMsgClaimRunnerRewards(creator string, amount string) *MsgClaimRunnerRewards {
	return &MsgClaimRunnerRewards{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgClaimRunnerRewards) Route() string {
	return RouterKey
}

func (msg *MsgClaimRunnerRewards) Type() string {
	return TypeMsgClaimRunnerRewards
}

func (msg *MsgClaimRunnerRewards) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimRunnerRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimRunnerRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimMotusRewards = "claim_motus_rewards"

var _ sdk.Msg = &MsgClaimMotusRewards{}

func NewMsgClaimMotusRewards(creator string, amount string) *MsgClaimMotusRewards {
	return &MsgClaimMotusRewards{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgClaimMotusRewards) Route() string {
	return RouterKey
}

func (msg *MsgClaimMotusRewards) Type() string {
	return TypeMsgClaimMotusRewards
}

func (msg *MsgClaimMotusRewards) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimMotusRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimMotusRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
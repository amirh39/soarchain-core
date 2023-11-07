package types

import (
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimChallengerRewards = "claim_challenger_rewards"

func NewMsgClaimChallengerRewards(creator string, amount string) *MsgClaimChallengerRewards {
	return &MsgClaimChallengerRewards{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgClaimChallengerRewards) Route() string {
	return RouterKey
}

func (msg *MsgClaimChallengerRewards) Type() string {
	return TypeMsgClaimChallengerRewards
}

func (msg *MsgClaimChallengerRewards) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimChallengerRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimChallengerRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "[NewMsgClaimChallengerRewards][ValidateBasic] failed. Invalid creator address: [ %s ] .", msg.Creator)
	}
	if isValidRewardAmount := utility.ValidateRewardAmount(msg.Amount); !isValidRewardAmount {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "[NewMsgClaimChallengerRewards][ValidateBasic] failed. Invalid reward amount: [ %s ].", msg.Amount)
	}
	return nil
}

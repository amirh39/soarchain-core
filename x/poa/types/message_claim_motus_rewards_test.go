package types

import (
	"testing"

	"soarchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgClaimMotusRewards_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgClaimMotusRewards
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgClaimMotusRewards{
				Creator: "invalid_address",
				Amount:  "1000udmotus",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgClaimMotusRewards{
				Creator: sample.AccAddress(),
				Amount:  "1000udmotus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

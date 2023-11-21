package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soar-robotics/soarchain-core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSelectRandomRunner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSelectRandomRunner
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSelectRandomRunner{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSelectRandomRunner{
				Creator: sample.AccAddress(),
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

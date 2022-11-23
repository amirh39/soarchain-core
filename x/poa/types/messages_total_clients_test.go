package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"soarchain/testutil/sample"
)

func TestMsgCreateTotalClients_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTotalClients
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTotalClients{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTotalClients{
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

func TestMsgUpdateTotalClients_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTotalClients
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTotalClients{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTotalClients{
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

func TestMsgDeleteTotalClients_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTotalClients
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTotalClients{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTotalClients{
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

package types

import (
	"testing"

	"soarchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRunnerChallenge_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRunnerChallenge
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRunnerChallenge{
				Creator:         "invalid_address",
				RunnerPubkey:    "runnerpubkey",
				ClientPubkeys:   []string{"clientpubkey"},
				ChallengeResult: "reward",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRunnerChallenge{
				Creator:         sample.AccAddress(),
				RunnerPubkey:    "runnerpubkey",
				ClientPubkeys:   []string{"clientpubkey"},
				ChallengeResult: "reward",
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

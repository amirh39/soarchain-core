package cli

import (
	"strconv"

	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRunnerChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner-challenge [runner-address] [v2n-device-type] [challenge-result]",
		Short: "Broadcast message runner-challenge",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRunnerAddress := args[0]
			argV2NDeviceType := args[1]
			argChallengeResult := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRunnerChallenge(
				clientCtx.GetFromAddress().String(),
				argRunnerAddress,
				argV2NDeviceType,
				argChallengeResult,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

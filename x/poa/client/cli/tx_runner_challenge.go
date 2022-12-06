package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"soarchain/x/poa/types"
)

var _ = strconv.Itoa(0)

func CmdRunnerChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner-challenge [runner-address] [challenge-result]",
		Short: "Broadcast message runner-challenge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRunnerAddress := args[0]
			argChallengeResult := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRunnerChallenge(
				clientCtx.GetFromAddress().String(),
				argRunnerAddress,
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

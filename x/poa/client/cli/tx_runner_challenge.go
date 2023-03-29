package cli

import (
	"strconv"

	"soarchain/x/poa/types"

	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRunnerChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner-challenge [runner-address] [v2nbx-device-array]",
		Short: "Broadcast message runner-challenge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRunnerAddress := args[0]
			argClientPubkeys := strings.Split(args[1], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRunnerChallenge(
				clientCtx.GetFromAddress().String(),
				argRunnerAddress,
				argClientPubkeys,
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

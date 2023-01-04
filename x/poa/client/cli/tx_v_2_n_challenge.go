package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"soarchain/x/poa/types"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdV2NChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "v-2-n-challenge [runner-address] [runner-result] [v-2-n-bx-address] [v-2-n-bx-result]",
		Short: "Broadcast message v2nChallenge",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRunnerAddress := args[0]
			argRunnerResult := args[1]
			argV2NBxAddress := strings.Split(args[2], listSeparator)
			argV2NBxResult := strings.Split(args[3], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgV2NChallenge(
				clientCtx.GetFromAddress().String(),
				argRunnerAddress,
				argRunnerResult,
				argV2NBxAddress,
				argV2NBxResult,
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

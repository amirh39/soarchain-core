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

func CmdV2VChallenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "v-2-v-challenge [rx-address] [rx-result] [bx-address] [bx-result]",
		Short: "Broadcast message v2vChallenge",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRxAddress := args[0]
			argRxResult := args[1]
			argBxAddress := strings.Split(args[2], listSeparator)
			argBxResult := strings.Split(args[3], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgV2VChallenge(
				clientCtx.GetFromAddress().String(),
				argRxAddress,
				argRxResult,
				argBxAddress,
				argBxResult,
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

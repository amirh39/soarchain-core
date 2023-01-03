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

func CmdChallengeService() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "challenge-service [challengee-address] [client-communication-mode] [challenge-result]",
		Short: "Broadcast message challenge-service",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argClientAddress := args[0]
			argClientCommunicationMode := args[1]
			argChallengeResult := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChallengeService(
				clientCtx.GetFromAddress().String(),
				argClientAddress,
				argClientCommunicationMode,
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

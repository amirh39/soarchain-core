package cli

import (
	"strconv"

	"github.com/amirh39/soarchain-core/x/poa/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdChallengeService() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "challenge-service [client-pubkey] [client-communication-mode] [challenge-result]",
		Short: "Broadcast message challenge-service",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argClientPubkey := args[0]
			argClientCommunicationMode := args[1]
			argChallengeResult := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChallengeService(
				clientCtx.GetFromAddress().String(),
				argClientPubkey,
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

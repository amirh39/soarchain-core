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

func CmdUnregisterClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unregister-client [pubkey] [fee]",
		Short: "Broadcast message unregister-client",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPubkey := args[0]
			argFee := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnregisterClient(
				clientCtx.GetFromAddress().String(),
				argPubkey,
				argFee,
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

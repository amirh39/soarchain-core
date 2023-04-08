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

func CmdGenClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-client [Certificate] [Signature]",
		Short: "Broadcast message gen-client",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argPubkey := "-----BEGIN CERTIFICATE-----\n" + args[0] + "\n-----END CERTIFICATE-----"
			argSignature := args[1] //TODO: ----- flag change

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenClient(
				clientCtx.GetFromAddress().String(),
				argPubkey,
				argSignature,
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

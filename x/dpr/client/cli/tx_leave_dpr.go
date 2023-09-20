package cli

import (
	"strconv"

	"soarchain/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdLeaveDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leave-dpr [pubkey] [pdrId]",
		Short: "Broadcast message leave-dpr",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argPubkey := args[0]
			argDprId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLeaveDpr(
				argPubkey,
				clientCtx.GetFromAddress().String(),
				argDprId,
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

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

func CmdDeactivateDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-dpr [pdrId] [duration]",
		Short: "Broadcast message activate-dpr",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argDprId := args[0]
			argDuration, _ := strconv.ParseUint(args[2], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeactivateDpr(
				clientCtx.GetFromAddress().String(),
				argDprId,
				argDuration,
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

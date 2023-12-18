package cli

import (
	"strconv"

	"github.com/amirh39/soarchain-core/x/dpr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-dpr [dprId][duration][MaxClientCount][DprBudget]",
		Short: "Broadcast message update-dpr",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			dprId := args[0]
			duration, _ := strconv.ParseUint(args[1], 10, 64)
			MaxClientCount, _ := strconv.ParseUint(args[2], 10, 64)
			DprBudget := args[3]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDpr(
				dprId,
				duration,
				MaxClientCount,
				DprBudget,
				clientCtx.GetFromAddress().String(),
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

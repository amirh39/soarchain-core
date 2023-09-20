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

func CmdUpdateDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-dpr [pdrId] [pidSupportedOneToTwnety] [pidSupportedTwentyOneToForthy] [pidSupportedForthyOneToSixty]",
		Short: "Broadcast message update-dpr",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			dprId := args[0]
			pidSupportedOneToTwnety, _ := strconv.ParseBool(args[1])
			pidSupportedTwentyOneToForthy, _ := strconv.ParseBool(args[2])
			pidSupportedForthyOneToSixty, _ := strconv.ParseBool(args[3])
			duration, _ := strconv.ParseUint(args[4], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDpr(
				dprId,
				pidSupportedOneToTwnety,
				pidSupportedTwentyOneToForthy,
				pidSupportedForthyOneToSixty,
				duration,
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

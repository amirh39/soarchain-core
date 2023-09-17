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

func CmdGenDpr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-dpr [pidSupportedOneToTwnety] [pidSupportedTwentyOneToForthy] [pidSupportedForthyOneToSixty] [lengthOfDpr]",
		Short: "Broadcast message gen-dpr",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			pidSupportedOneToTwnety, _ := strconv.ParseBool(args[0])
			pidSupportedTwentyOneToForthy, _ := strconv.ParseBool(args[1])
			pidSupportedForthyOneToSixty, _ := strconv.ParseBool(args[2])
			lengthOfDpr, _ := strconv.ParseUint(args[3], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenDpr(
				clientCtx.GetFromAddress().String(),
				pidSupportedOneToTwnety,
				pidSupportedTwentyOneToForthy,
				pidSupportedForthyOneToSixty,
				lengthOfDpr,
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
